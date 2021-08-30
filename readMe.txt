Используется две api

1. Mapbox: (URL: https://api.mapbox.com/geocoding/v5/mapbox.places/London.json?access_token='my_access_token'&limit=1)
    a. Вводим название города
    б. Получаем latitude(ширину), longitude(долготу), location(локацию)

2.Weatherstack:(URL: http://api.weatherstack.com/current?access_key='my_access_key'&query=latitude,longitude)
    a. Получает координаты из api Mapbox (latitude,longitude)
    b. Возвращает temperature(температуру), weather_descriptions(описание погоды), Humidity(влажность)


Собираем в Json:
    {
        "city":"London",
        "latitude":"51.517",
        "longitude":"-0.106"
        "temperature":"13"
        "weather_descriptions":["sunny"]
        "Humidity":13%
        "data":2021:08:21 10:50
    }

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

{
    name: "London"
    full_name: "то что писал пользователь"

}
