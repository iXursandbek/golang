package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("Maxsulot topilmadi")
	ErrCantDecodeProducts = errors.New("Maxsulot topilmadi")
	ErrUserIdIsNotValid   = errors.New("Foydalanuvchi yaroqsiz")
	ErrCantUpdateUser     = errors.New("Bu maxsulotni savatchaga qoshib bolmaydi")
	ErrCantRemoveItemCart = errors.New("Mu elemntni savatchadan olib tashlab bo'lmaydi")
	ErrCantGetItem        = errors.New("Maxsulotni savatdan olib bo'lmadi")
	ErrCantBuyCartItem    = errors.New("Xaridni yangilay olmaydi")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
