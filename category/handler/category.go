package handler

import (
	"context"
	"git.imocc.com/cap1573/category/common"
	"git.imocc.com/cap1573/category/domain/model"
	"git.imocc.com/cap1573/category/domain/service"
	category "git.imocc.com/cap1573/category/proto"
	"go-micro.dev/v4/util/log"
)

type Category struct{

	CategoryDateService service.ICategoryDataService
}

// 提供创建分类服务
func (c *Category) CreateCategory(ctx context.Context,request *category.CategoryRequest,response *category.CreateCategoryResponse) error  {
	category := &model.Category{}
	// 赋值
	err := common.SwapTo(request,category)
	if  err  != nil{
		return  err
	}

	categoryId , err  := c.CategoryDateService.AddCategory(category)
	if  err  != nil{
		return  err
	}
	response.Message ="分类添加成功"
	response.CategoryId =categoryId
	return  nil
}

// 提供服务更新服务
 func  (c  *Category) UpdateCategory (ctx  context.Context ,request *category.CategoryRequest,response *category.UpdateCategoryResponse) error{
 	category  :=model.Category{}

 	err  := common.SwapTo(response , category)

 	if  err  != nil{
 		return  err
	}
	err  =  c.CategoryDateService.UpdateCategory(&category)
	if  err  != nil{
		return  err
	}

	 response.Message = "分类更新成功！！！"
	return nil
 }
// 提供分类删除服务
func (c *Category)DeleteCategory(ctx context.Context, request *category.DeleteCategoryRequest,response *category.DeleteCatehoryResponse) error{
	err  := c.CategoryDateService.DeleteCategory(request.CategoryId)
	if  err != nil{
		return  nil
	}

	response.Message ="删除成功"
	return nil
}
// 根据分类名称查找分类
func  (c  *Category) FindCategoryByName(ctx context.Context, request *category.FindByNameRequest, response *category.CategoryResponse) error{
	category, err  := c.CategoryDateService.FindCategoryByName(request.CategoryName)
	if  err != nil{
		return  err
	}
	return  common.SwapTo(category,response)
}

// 根据分类ID查询分类
func (c *Category)FindCategoryByID(ctx context.Context, request *category.FindByIdRequest, response *category.CategoryResponse) error{
	category , err  := c.CategoryDateService.FindCategoryByID(request.CategoryId)
	if  err  != nil{
		return  err
	}

	return  common.SwapTo(category,response)
}

// 根据层级查找分类信息
func (c *Category) FindCategoryByLevel(ctx context.Context,  request *category.FindByLevelRequest , response  *category.FindAllResponse)  error{
	categorySlice  , err  := c.CategoryDateService.FindCategoryByLevel(request.Level)
	if  err != nil{
		return err
	}
	 categoryToResponse(categorySlice,response)
	return  nil
}

func  categoryToResponse (categorySlice []model.Category, response *category.FindAllResponse){
	for  _, cg := range categorySlice {
		cr := &category.CategoryResponse{}
		err  := common.SwapTo(cg, cr)
		if  err  != nil{
			log.Error(err)
			break
		}
		response.Category  = append(response.Category,cr)

	}
}
// 根据Parent查找分类信息
func (c *Category) FindCategoryByParent(ctx context.Context, request *category.FindByParentRequest, response *category.FindAllResponse)  error{
  categorySlice  , err :=	c.CategoryDateService.FindCategoryByParent(request.PrentId)
  if  err != nil{
  	return  err
  }
  categoryToResponse(categorySlice,response)
  return nil
}

func (c *Category)  FidAllCategory(ctx context.Context, request *category.FindAllRequest ,response *category.FindAllResponse) error{

	categorySlice  , err :=	c.CategoryDateService.FindAllCategory()
	if  err != nil{
		return  err
	}
	categoryToResponse(categorySlice,response)
	return  nil
}
