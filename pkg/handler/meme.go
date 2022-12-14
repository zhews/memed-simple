package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	memeConfig "github.com/zhews/memed-simple/pkg/config/meme"
	"github.com/zhews/memed-simple/pkg/handler/dto"
	"github.com/zhews/memed-simple/pkg/service"
	"log"
	"net/http"
)

type MemeHandler struct {
	Config  memeConfig.Config
	Service service.MemeService
}

func (mh *MemeHandler) HandleGetMemes(ctx *fiber.Ctx) error {
	memes, err := mh.Service.GetMemes()
	if err != nil {
		log.Println("Could not get memes: ", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	dtoMemes := make([]dto.MemeResponse, 0)
	creatorCache := map[uuid.UUID]dto.UserResponse{}
	for _, meme := range memes {
		var creator dto.UserResponse
		creator, ok := creatorCache[meme.CreatedBy]
		if !ok {
			response, err := http.Get(fmt.Sprintf("%s%s/%s", mh.Config.UserMicroservice, mh.Config.UserEndpoint, meme.CreatedBy))
			if err != nil {
				log.Println("Could not get user information: ", err)
				return ctx.SendStatus(fiber.StatusInternalServerError)
			}
			var newCreator dto.UserResponse
			err = json.NewDecoder(response.Body).Decode(&creator)
			if err != nil {
				log.Println("Could not decode user information: ", err)
				return ctx.SendStatus(fiber.StatusInternalServerError)
			}
			creatorCache[meme.CreatedBy] = newCreator
			err = response.Body.Close()
			if err != nil {
				log.Println("Could not close response body:", err)
			}
		}
		dtoMeme := dto.MemeResponse{
			Id:        meme.Id,
			Title:     meme.Title,
			Image:     fmt.Sprintf("/image/%s", meme.Image),
			Creator:   creator,
			CreatedAt: meme.CreatedAt,
			UpdatedAt: meme.UpdatedAt,
		}
		dtoMemes = append(dtoMemes, dtoMeme)
	}
	return ctx.Status(fiber.StatusOK).JSON(dtoMemes)
}

func (mh *MemeHandler) HandleGetMeme(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	meme, err := mh.Service.GetMemeById(id)
	if err != nil {
		log.Println("Could not get meme by id: ", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	response, err := http.Get(fmt.Sprintf("%s%s/%s", mh.Config.UserMicroservice, mh.Config.UserEndpoint, meme.CreatedBy))
	if err != nil {
		log.Println("Could not get user information: ", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	var creator dto.UserResponse
	err = json.NewDecoder(response.Body).Decode(&creator)
	if err != nil {
		log.Println("Could not decode user information: ", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	dtoMeme := dto.MemeResponse{
		Id:        meme.Id,
		Title:     meme.Title,
		Image:     fmt.Sprintf("/image/%s", meme.Image),
		Creator:   creator,
		CreatedAt: meme.CreatedAt,
		UpdatedAt: meme.UpdatedAt,
	}
	err = response.Body.Close()
	if err != nil {
		log.Println("Could not close response body:", err)
	}
	return ctx.Status(fiber.StatusOK).JSON(dtoMeme)
}

func (mh *MemeHandler) HandleUploadMeme(ctx *fiber.Ctx) error {
	title := ctx.FormValue("title")
	meme, err := ctx.FormFile("meme")
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	contentType := meme.Header.Get("Content-Type")
	memeFile, err := meme.Open()
	token, ok := ctx.Locals("user").(*jwt.Token)
	if !ok {
		log.Println("Could not parse token from context: ", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	claims := token.Claims.(jwt.MapClaims)
	userIdString := claims["sub"].(string)
	userId, err := uuid.Parse(userIdString)
	if err != nil {
		log.Println("Could not parse uuid from token: ", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	err = mh.Service.UploadMeme(title, contentType, memeFile, userId)
	if err != nil {
		log.Println("Could not upload file:", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (mh *MemeHandler) HandleUpdateMeme(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	var request dto.UpdateMemeRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	token, ok := ctx.Locals("user").(*jwt.Token)
	if !ok {
		log.Println("Could not parse token from context: ", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	claims := token.Claims.(jwt.MapClaims)
	userIdString := claims["sub"].(string)
	userId, err := uuid.Parse(userIdString)
	if err != nil {
		log.Println("Could not parse uuid from claims: ", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	err = mh.Service.UpdateMeme(id, request.Title, userId)
	if err != nil {
		log.Println("Could not update meme: ", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (mh *MemeHandler) HandleDeleteMeme(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	err = mh.Service.DeleteMeme(id)
	if err != nil {
		log.Println("Could not delete meme: ", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
