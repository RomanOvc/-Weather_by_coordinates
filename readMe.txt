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
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


Redis.
Реализовать хранение access токена в редисе и его проверку в мидлваре.

username -> jwtToken
(ex +) exRedis = exToken

login -> set token to redis
middleware -> validate -> get token.username from redis  tokenRedis == Token
