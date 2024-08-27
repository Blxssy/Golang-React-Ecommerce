package service

import (
	"github.com/Blxssy/Golang-React-Ecommerce/internal/container"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/models"
	"github.com/Blxssy/Golang-React-Ecommerce/internal/storage"
)

type CartService interface {
	AddItem(userID uint, productID uint, quantity int) error
	RemoveItem(userID uint, productID uint) error
	UpdateItemQuantity(userID uint, productID uint, quantity int) error
	GetCart(userID uint) (*models.Cart, error)
	ClearCart(userID uint) error
}

type cartService struct {
	container container.Container
}

func NewCartService(container container.Container) CartService {
	return &cartService{
		container: container,
	}
}

func (s *cartService) AddItem(userID uint, productID uint, quantity int) error {
	//logger := s.container.GetLogger()

	// Загрузка корзины вместе с элементами
	var cart models.Cart
	if err := s.container.GetRepository().Preload("Items").First(&cart, "user_id = ?", userID).Error; err != nil {
		return err
	}

	//logger.Info("cartInfo", slog.Any("cart", cart))

	var cartItem *models.CartItem
	for i := range cart.Items {
		if cart.Items[i].ProductId == productID {
			cartItem = &cart.Items[i]
			break
		}
	}

	if cartItem != nil {
		// Если предмет найден, увеличиваем его количество
		cartItem.Quantity += quantity
		if err := s.container.GetRepository().Save(cartItem).Error; err != nil {
			return err
		}
	} else {
		// Если предмет не найден, создаем новый CartItem
		newCartItem := models.CartItem{
			CartID:    cart.ID,
			ProductId: productID,
			Quantity:  quantity,
		}

		// Добавляем новый элемент в корзину с помощью Association
		if err := s.container.GetRepository().Model(&cart).Association("Items").Append(&newCartItem); err != nil {
			return err
		}
	}

	// Сохранение корзины (необязательно, если изменения в элементах автоматически сохраняются)
	if err := s.SaveCart(&cart); err != nil {
		return err
	}

	//logger.Info("Item added to cart", slog.Any("cart", cart))
	return nil
}

func (s *cartService) RemoveItem(userID uint, productID uint) error {
	return nil
}

func (s *cartService) UpdateItemQuantity(userID uint, productID uint, quantity int) error {
	return nil
}

func (s *cartService) GetCart(userID uint) (*models.Cart, error) {
	var cart models.Cart

	// Загрузка корзины вместе с элементами
	if err := s.container.GetRepository().Preload("Items").First(&cart, userID).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func (s *cartService) ClearCart(userID uint) error {
	return nil
}

func (s *cartService) SaveCart(cart *models.Cart) error {
	repo := s.container.GetRepository()
	var err error

	if txerr := repo.Transaction(func(tx storage.Storage) error {
		_, err = txSaveCart(tx, cart)
		return err
	}); txerr != nil {
		s.container.GetLogger().Error(txerr.Error())
		return txerr
	}

	return nil
}

func txSaveCart(tx storage.Storage, cart *models.Cart) (*models.Cart, error) {
	var err error

	if err = cart.Save(tx); err != nil {
		return nil, err
	}

	return cart, nil
}
