package routes

import (
	"Basketball_academy/controllers"
	"Basketball_academy/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")

	// Пользователи
	userCtrl := controllers.NewCRUDController[models.User](db)
	users := api.Group("/users")
	{
		users.GET("", userCtrl.GetAll)
		users.GET("/:id", userCtrl.GetByID)
		users.POST("", userCtrl.Create)
		users.PUT("/:id", userCtrl.Update)
		users.DELETE("/:id", userCtrl.Delete)
	}

	// Сотрудники академии
	academyStaffCtrl := controllers.NewCRUDController[models.AcademyStaff](db)
	academyStaff := api.Group("/academy-staff")
	{
		academyStaff.GET("", academyStaffCtrl.GetAll)
		academyStaff.GET("/:id", academyStaffCtrl.GetByID)
		academyStaff.POST("", academyStaffCtrl.Create)
		academyStaff.PUT("/:id", academyStaffCtrl.Update)
		academyStaff.DELETE("/:id", academyStaffCtrl.Delete)
	}

	// Тренеры
	trainerCtrl := controllers.NewCRUDController[models.Trainer](db)
	trainers := api.Group("/trainers")
	{
		trainers.GET("", trainerCtrl.GetAll)
		trainers.GET("/:id", trainerCtrl.GetByID)
		trainers.POST("", trainerCtrl.Create)
		trainers.PUT("/:id", trainerCtrl.Update)
		trainers.DELETE("/:id", trainerCtrl.Delete)
	}

	// Чаты
	chatCtrl := controllers.NewCRUDController[models.Chat](db)
	chats := api.Group("/chats")
	{
		chats.GET("", chatCtrl.GetAll)
		chats.GET("/:id", chatCtrl.GetByID)
		chats.POST("", chatCtrl.Create)
		chats.PUT("/:id", chatCtrl.Update)
		chats.DELETE("/:id", chatCtrl.Delete)
	}

	// Сообщения в чате
	chatMsgCtrl := controllers.NewCRUDController[models.ChatMessage](db)
	chatMessages := api.Group("/chat-messages")
	{
		chatMessages.GET("", chatMsgCtrl.GetAll)
		chatMessages.GET("/:id", chatMsgCtrl.GetByID)
		chatMessages.POST("", chatMsgCtrl.Create)
		chatMessages.PUT("/:id", chatMsgCtrl.Update)
		chatMessages.DELETE("/:id", chatMsgCtrl.Delete)
	}

	// Новости
	newsCtrl := controllers.NewCRUDController[models.News](db)
	news := api.Group("/news")
	{
		news.GET("", newsCtrl.GetAll)
		news.GET("/:id", newsCtrl.GetByID)
		news.POST("", newsCtrl.Create)
		news.PUT("/:id", newsCtrl.Update)
		news.DELETE("/:id", newsCtrl.Delete)
	}

	// Позиции заказа (OrderItem)
	orderItemCtrl := controllers.NewCRUDController[models.OrderItem](db)
	orderItems := api.Group("/order-items")
	{
		orderItems.GET("", orderItemCtrl.GetAll)
		orderItems.GET("/:id", orderItemCtrl.GetByID)
		orderItems.POST("", orderItemCtrl.Create)
		orderItems.PUT("/:id", orderItemCtrl.Update)
		orderItems.DELETE("/:id", orderItemCtrl.Delete)
	}

	// Заказы
	orderCtrl := controllers.NewCRUDController[models.Order](db)
	orders := api.Group("/orders")
	{
		orders.GET("", orderCtrl.GetAll)
		orders.GET("/:id", orderCtrl.GetByID)
		orders.POST("", orderCtrl.Create)
		orders.PUT("/:id", orderCtrl.Update)
		orders.DELETE("/:id", orderCtrl.Delete)
	}

	// Платежи
	paymentCtrl := controllers.NewCRUDController[models.Payment](db)
	payments := api.Group("/payments")
	{
		payments.GET("", paymentCtrl.GetAll)
		payments.GET("/:id", paymentCtrl.GetByID)
		payments.POST("", paymentCtrl.Create)
		payments.PUT("/:id", paymentCtrl.Update)
		payments.DELETE("/:id", paymentCtrl.Delete)
	}

	// Товары
	productCtrl := controllers.NewCRUDController[models.Product](db)
	products := api.Group("/products")
	{
		products.GET("", productCtrl.GetAll)
		products.GET("/:id", productCtrl.GetByID)
		products.POST("", productCtrl.Create)
		products.PUT("/:id", productCtrl.Update)
		products.DELETE("/:id", productCtrl.Delete)
	}

	// Промокоды
	promoCodeCtrl := controllers.NewCRUDController[models.PromoCode](db)
	promoCodes := api.Group("/promo-codes")
	{
		promoCodes.GET("", promoCodeCtrl.GetAll)
		promoCodes.GET("/:id", promoCodeCtrl.GetByID)
		promoCodes.POST("", promoCodeCtrl.Create)
		promoCodes.PUT("/:id", promoCodeCtrl.Update)
		promoCodes.DELETE("/:id", promoCodeCtrl.Delete)
	}

	// Акции (Promotion)
	promotionCtrl := controllers.NewCRUDController[models.Promotion](db)
	promotions := api.Group("/promotions")
	{
		promotions.GET("", promotionCtrl.GetAll)
		promotions.GET("/:id", promotionCtrl.GetByID)
		promotions.POST("", promotionCtrl.Create)
		promotions.PUT("/:id", promotionCtrl.Update)
		promotions.DELETE("/:id", promotionCtrl.Delete)
	}

	// Статические страницы
	staticPageCtrl := controllers.NewCRUDController[models.StaticPage](db)
	staticPages := api.Group("/static-pages")
	{
		staticPages.GET("", staticPageCtrl.GetAll)
		staticPages.GET("/:id", staticPageCtrl.GetByID)
		staticPages.POST("", staticPageCtrl.Create)
		staticPages.PUT("/:id", staticPageCtrl.Update)
		staticPages.DELETE("/:id", staticPageCtrl.Delete)
	}

	// Техподдержка (SupportTicket)
	supportTicketCtrl := controllers.NewCRUDController[models.SupportTicket](db)
	supportTickets := api.Group("/support-tickets")
	{
		supportTickets.GET("", supportTicketCtrl.GetAll)
		supportTickets.GET("/:id", supportTicketCtrl.GetByID)
		supportTickets.POST("", supportTicketCtrl.Create)
		supportTickets.PUT("/:id", supportTicketCtrl.Update)
		supportTickets.DELETE("/:id", supportTicketCtrl.Delete)
	}

	// Группы тренировок
	trainingGroupCtrl := controllers.NewCRUDController[models.TrainingGroup](db)
	trainingGroups := api.Group("/training-groups")
	{
		trainingGroups.GET("", trainingGroupCtrl.GetAll)
		trainingGroups.GET("/:id", trainingGroupCtrl.GetByID)
		trainingGroups.POST("", trainingGroupCtrl.Create)
		trainingGroups.PUT("/:id", trainingGroupCtrl.Update)
		trainingGroups.DELETE("/:id", trainingGroupCtrl.Delete)
	}

	// Регистрации на тренировки
	trainingRegistrationCtrl := controllers.NewCRUDController[models.TrainingRegistration](db)
	trainingRegistrations := api.Group("/training-registrations")
	{
		trainingRegistrations.GET("", trainingRegistrationCtrl.GetAll)
		trainingRegistrations.GET("/:id", trainingRegistrationCtrl.GetByID)
		trainingRegistrations.POST("", trainingRegistrationCtrl.Create)
		trainingRegistrations.PUT("/:id", trainingRegistrationCtrl.Update)
		trainingRegistrations.DELETE("/:id", trainingRegistrationCtrl.Delete)
	}

	// Тренировки (TrainingSession)
	trainingSessionCtrl := controllers.NewCRUDController[models.TrainingSession](db)
	trainingSessions := api.Group("/training-sessions")
	{
		trainingSessions.GET("", trainingSessionCtrl.GetAll)
		trainingSessions.GET("/:id", trainingSessionCtrl.GetByID)
		trainingSessions.POST("", trainingSessionCtrl.Create)
		trainingSessions.PUT("/:id", trainingSessionCtrl.Update)
		trainingSessions.DELETE("/:id", trainingSessionCtrl.Delete)
	}

	// Мои тренировки (MyTraining)
	myTrainingCtrl := controllers.NewCRUDController[models.MyTraining](db)
	myTrainings := api.Group("/my-trainings")
	{
		myTrainings.GET("", myTrainingCtrl.GetAll)
		myTrainings.GET("/:id", myTrainingCtrl.GetByID)
		myTrainings.POST("", myTrainingCtrl.Create)
		myTrainings.PUT("/:id", myTrainingCtrl.Update)
		myTrainings.DELETE("/:id", myTrainingCtrl.Delete)
	}

	// Уведомления
	notificationCtrl := controllers.NewCRUDController[models.Notification](db)
	notifications := api.Group("/notifications")
	{
		notifications.GET("", notificationCtrl.GetAll)
		notifications.GET("/:id", notificationCtrl.GetByID)
		notifications.POST("", notificationCtrl.Create)
		notifications.PUT("/:id", notificationCtrl.Update)
		notifications.DELETE("/:id", notificationCtrl.Delete)
	}

	// Клубы
	clubCtrl := controllers.NewCRUDController[models.Club](db)
	clubs := api.Group("/clubs")
	{
		clubs.GET("", clubCtrl.GetAll)
		clubs.GET("/:id", clubCtrl.GetByID)
		clubs.POST("", clubCtrl.Create)
		clubs.PUT("/:id", clubCtrl.Update)
		clubs.DELETE("/:id", clubCtrl.Delete)
	}

	// Корзины
	cartCtrl := controllers.NewCRUDController[models.Cart](db)
	carts := api.Group("/carts")
	{
		carts.GET("", cartCtrl.GetAll)
		carts.GET("/:id", cartCtrl.GetByID)
		carts.POST("", cartCtrl.Create)
		carts.PUT("/:id", cartCtrl.Update)
		carts.DELETE("/:id", cartCtrl.Delete)
	}

	// Товары в корзине (CartItem)
	cartItemCtrl := controllers.NewCRUDController[models.CartItem](db)
	cartItems := api.Group("/cart-items")
	{
		cartItems.GET("", cartItemCtrl.GetAll)
		cartItems.GET("/:id", cartItemCtrl.GetByID)
		cartItems.POST("", cartItemCtrl.Create)
		cartItems.PUT("/:id", cartItemCtrl.Update)
		cartItems.DELETE("/:id", cartItemCtrl.Delete)
	}
}
