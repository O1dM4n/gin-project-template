package cart

import (
	"gorm.io/gorm"
	"oldman.study/gin-project-template/domain/product"
	"oldman.study/gin-project-template/domain/user"
)

// 购物车结构体
type Cart struct {
	gorm.Model
	UserID uint
	User   user.User `gorm:"foreignKey:ID;references:UserID"`
}

// 实例化
func NewCart(uid uint) *Cart {
	return &Cart{
		UserID: uid,
	}
}

// Item
type Item struct {
	gorm.Model
	Product   product.Product `gorm:"foreignKey:ProductID"`
	ProductID uint
	Count     int
	CartID    uint
	Cart      Cart `gorm:"foreignKey:CartID" json:"-"`
}

// 创建Item
func NewCartItem(productId uint, cartId uint, count int) *Item {
	return &Item{
		ProductID: productId,
		Count:     count,
		CartID:    cartId,
	}
}
