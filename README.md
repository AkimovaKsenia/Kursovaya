Cinema mastery - CRM система для работников справочной службы кинотеатров города.
![img.png](img/img.png)
## Установка и запуск
### 1 способ. Клонирование репозитория
1. Склонируйте репозиторий
```shell
git clone https://github.com/AkimovaKsenia/Kursovaya.git
```
2. Запустите docker-compose
```shell
docker compose up -d
```
### 2 способ. Без клонирования всего репозитория
1. Сохраните себе файл **docker-compose.server.yml**
2. Запустите его
```shell
docker compose -f .\docker-compose.server.yml up -d
```