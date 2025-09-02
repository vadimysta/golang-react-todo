# ✅ Todo App - Fullstack Go + React

Повнофункціональний додаток для управління задачами з повним стеком технологій: **Go backend + React frontend + MongoDB**.

## 🚀 Особливості

### Backend (Go)
- **RESTful API** з CRUD операціями
- **Fiber framework** для високопродуктивних HTTP запитів
- **MongoDB** з офіційним драйвером
- **CORS middleware** для кросс-доменних запитів
- **Структуроване логування** з slog

### Frontend (React)
- **Сучасний інтерфейс** з українською локалізацією
- **React Hooks** для управління станом
- **Vite** для швидкої збірки
- **Адаптивний дизайн** з CSS Flexbox/Grid
- **Оптимістичні оновлення** UI

### База даних
- **MongoDB** з коллекцією `todos`
- **Автогенерація ObjectID** для документів
- **Масштабована архітектура** для зберігання даних

## 📦 Встановлення та запуск

### 1. Клонування репозиторію
```bash
git clone https://github.com/vadimysta/golang-react-todo
cd golang-react-todo

### запуск сервера
cd backend 
go run cmd/api/main.go

### запуск UI
cd frontend
npm run dev