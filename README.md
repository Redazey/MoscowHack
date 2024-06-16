# Хакатон Лидеры цифровой трансформации 2024
Наш проект представляет собой web-приложение, которое предоставляет пользователям функционал коммуникационной площадки. Эта платформа позволяет размещать новости, обмениваться комментариями, получать уведомления и автоматизировать ключевые этапы процесса рекрутинга.

# Технологический стек
## Backend
- Golang
- gRPC
## Frontend
- HTML
- SCSS
- TypeScript
- React
## Базы данных
- PostgreSQL
- Redis (in-memory)

# Основные функции
## Коммуникационная платформа
- Размещение новостей.
- Обмен комментариями.
- Получение уведомлений.
## Рекрутинговые функции
- Интеграция рабочего календаря рекрутера.
- Бронирование времени.
- Настройка количества интервью в день.
- Настройка перерывов между интервью.
- Напоминания об интервью соискателям и рекрутерам.
- Возможность поделиться страницей соискателя в социальных сетях.
- Автоматическое приветствие в чате с кандидатом.
- Генерация названий встреч.
- Создание конференций в Zoom/Google Meet/Telemost при назначении встречи.

## Документация
Пользовательскую документацию можно получить по данной [ссылке](https://redazey.github.io/MoscowHack/).

# Установка и запуск
- Клонируйте репозиторий:
```
git clone https://github.com/Redazey/MoscowHack.git
```
- Перейдите в директорию проекта:
```
cd MoscowHack
```
- Установите зависимости для Backend и Frontend:
```
cd backend
go mod download
cd ../frontend
npm install
```
- Запустите Backend:
```
cd ../backend
go run main.go
```
- Запустите Frontend:
```
cd frontend
npm install
npm run dev
```

# Вкладчики
- Косолапов Кирилл - Fullstack-разработчик
- Аксенов Илья - Backend-разработчик
- Нешкреба Кирилл - Fullstack-разработчик
- Назаренко Иван - Frontend-разработчик

# Лицензия
Этот проект лицензируется под лицензией MIT. Подробнее см. [LICENSE](https://github.com/Redazey/MoscowHack/blob/main/LICENSE).
