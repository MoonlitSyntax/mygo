// 文件: mygo/wire/wire.go
//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"mygo/internal/db"

	"mygo/internal/controller"
	"mygo/internal/repository"
	"mygo/internal/service"
)

var ArticleSet = wire.NewSet(
	repository.NewArticleRepository, // -> ArticleRepository
	service.NewArticleService,       // -> ArticleService
	controller.NewArticleController, // -> *ArticleController
)

var CategorySet = wire.NewSet(
	repository.NewCategoryRepository,
	service.NewCategoryService,
	controller.NewCategoryController,
)

var TagSet = wire.NewSet(
	repository.NewTagRepository,
	service.NewTagService,
	controller.NewTagController,
)

var UserSet = wire.NewSet(
	repository.NewUserRepository,
	service.NewUserService,
	controller.NewUserController,
)

var MetaDataSet = wire.NewSet(
	repository.NewMetaDataRepository,
	service.NewMetaDataService,
	controller.NewMetaDataController,
)

var ProviderSetAll = wire.NewSet(
	ArticleSet,
	CategorySet,
	TagSet,
	UserSet,
	MetaDataSet,
)

type AppControllers struct {
	ArticleCtrl  *controller.ArticleController
	CategoryCtrl *controller.CategoryController
	TagCtrl      *controller.TagController
	UserCtrl     *controller.UserController
	MetaDataCtrl *controller.MetaDataController
}

func InitAppControllers() *AppControllers {
	wire.Build(
		db.GetDB,
		ProviderSetAll,

		wire.Struct(new(AppControllers),
			"ArticleCtrl",
			"CategoryCtrl",
			"TagCtrl",
			"UserCtrl",
			"MetaDataCtrl",
		),
	)
	return nil
}
