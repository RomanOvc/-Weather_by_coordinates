Используется две api

1. Mapbox: (URL: https://api.mapbox.com/geocoding/v5/mapbox.places/London.json?access_token='my_access_token'&limit=1)
    a. Вводим название города
    б. Получаем latitude(ширину), longitude(долготу), location(локацию)

2.Weatherstack:(URL: http://api.weatherstack.com/current?access_key='my_access_key'&query=latitude,longitude)
    a. Получает координаты из api Mapbox (latitude,longitude)
    b. Возвращает temperature(температуру), weather_descriptions(описание погоды), Humidity(влажность)

3. Собираем в Json:
    {
        "data_id":"1"
        "request":"london"
	    "city":"London"
	    "latitude":"51.507321899999994"
	    "longitude":"-0.12764739999999997"
	    "temperateure":"13"
	    "weather_descriptions":"[sunny]"
	    "humidity":"13"
	    "data":"2021-09-13"
    }

4. Если данных нет в базе, то получаем запрос от api и записываем данные в бд. Если есть такой запрос в базе выдать ответ.

5. Сделать Аутентификацию по Jwt.  В Jwt передать user_id, username, exp(время жизни токена). База данных выыглядит следующим образом (смотреть в sql.txt)
Сделать два эндпоинта (вернуть все запросы по user_id(user_id берется из токена) и добавление данных) и обернуть их миделваром.  
Структура Json  
   {
        "data_id":"1",
        "request":"london",
	    "city":"London",
	    "latitude":"51.507321899999994",
	    "longitude":"-0.12764739999999997",
	    "temperateure":"13",
	    "weather_descriptions":"[sunny]",
	    "humidity":"13",
        "user_id":"1",
	    "data":"2021-09-13"
    }

6. Добавить созданный токен в Redis.