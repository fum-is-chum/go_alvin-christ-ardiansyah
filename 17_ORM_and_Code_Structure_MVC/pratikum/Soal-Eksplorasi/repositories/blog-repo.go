package repositories

import (
	"errors"
	"soal-eksplorasi/config"
	"soal-eksplorasi/models"
)

func SelectBlogs() ([]models.Blog, error) {
	var blogs []models.Blog

	tx := config.DB.Order("created_at desc").Find(&blogs)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return blogs, nil
}

func SelectBlogById(Id uint) (models.Blog, error) {
	var blog models.Blog
	tx := config.DB.First(&blog, Id)
	if tx.Error != nil {
		return models.Blog{}, tx.Error
	}

	return blog, nil
}

func InsertBlog(blog models.Blog) error {
	var user models.User

	tx := config.DB.First(&user, blog.UserID)
	if tx.Error != nil {
		return errors.New("UserId Invalid or Not Found")
	}

	tx = config.DB.Create(&blog)
	if tx.Error != nil {
		return errors.New("Insert blog failed")
	}
	return nil
}

func DeleteBlog(Id uint) error {
	tx := config.DB.Delete(&models.Blog{}, Id)
	if tx.Error != nil {
		return errors.New("Delete blog failed")
	}
	return nil
}

func UpdateBlog(blog models.Blog) error {
	var user models.User

	tx := config.DB.First(&user, blog.UserID)
	if tx.Error != nil {
		return errors.New("UserId Invalid or Not Found")
	}
	
	updatedBlog := models.Blog{
		Title: blog.Title,
		Content: blog.Content,
		UserID: blog.UserID,
	}

	tx = config.DB.Model(&models.Blog{}).Where("id = ?", blog.ID).Updates(updatedBlog)
	if tx.Error != nil {
		return errors.New("Update blog failed")
	}

	if tx.RowsAffected == 0 {
		return errors.New("Blog not found")
	}

	return nil
}
