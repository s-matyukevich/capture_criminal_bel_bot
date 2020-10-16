package config

import (
	"fmt"
	"strings"

	"gitlab.com/opennota/morph"
)

var LabelKeywords = map[string]map[string]interface{}{
	"alien": {"тихарь": nil, "тихорь": nil, "тихори": nil, "космонавт": nil, "лошка": nil, "петушки": nil,
		"оливка": nil, "шакал": nil},
	"police": {"мент": nil, "мусор": nil, "омон": nil},
	"policeCar": {"бус": nil, "бусик": nil, "бусы": nil, "буса": nil, "автозак": nil, "автобус": nil, "минивен": nil, "машина": nil, "бобик": nil,
		"газель": nil, "газел": nil, "гаи": nil, "микроавтобус": nil, "микробус": nil, "форд": nil, "тонированный": nil, "маз": nil},
}

var LocationMap = map[string]interface{}{}

func init() {
	if err := morph.Init(); err != nil {
		panic(err)
	}

	locations := strings.Split(GeoLocations, "\n")
	for _, l := range locations {
		l = strings.ToLower(l)

		words := strings.Split(l, " ")
		norm := []string{}
		for _, w := range words {
			_, norms, _ := morph.Parse(w)
			if len(norms) > 0 {
				norm = append(norm, norms[0])
			} else {
				norm = append(norm, w)
			}
		}
		res := strings.Join(norm, " ")
		fmt.Printf("%s => %v\n", l, res)
		LocationMap[res] = l
	}
}

const GeoLocations = `Абрикосовая
Авакяна
Авангардная
Авангардный
Авиации
Авроровская
Автодоровская
Автодоровский
Автозаводская
Автозаводской
Автомобилистов
Агатовый
Шемеша
Азгура
Азизова
Азовская
Айвазовского
Айвазовского
Высоцкого
Вышелесского
Вышелесского
Жебрака
Жебрака
Карского
Красина
Купревича
Курчатова
Фёдорова
Академическая
Аладовых
Александровская
Александровский сквер
Бачило
Гаруна
Дудара
Дудара
Алибегова
Аллейная
Пашкевич
Алтайская
Алтайский
Амбулаторная
Амбулаторный
Амураторская
Амурская
Амурский
Ангарская
Ангарский
Андреевская
Аннаева
Антоненко
Антоновская
Аполинарьевская
Аранская
Смолича
Смолича
Арктическая 
Артиллеристов
Архитектора Заборского
Асаналиева
Асаналиева
Ауэзова
Аэродромная
Аэродромный
Аэрофлотская
Аэрофлотский 
Бабушкина
Багратиона
Багратиона 
Багратиона 
Багряная
Багрянцева
Багряный
Базисная 
Базисная 
Базисный
Байкальская
Бакинская
Балтийская
Бангалор
Барамзиной
Барановщина
Басиаловский 
Басиаловский 
Баторинская
Баторинский
Бегомльская
Белецкого
Белинского
Беломорская
Белорусская
Бельского
Бельчицкая
Беляева
Березинская
Берёзогорская
Берестянская
Берсона
Бетонный
Бехтерева
Бехтерева
Библиотечная
Бирюзова
Бобруйская
Берута
Болотная
Болотникова
Болотникова
Больничный
Большая Слепня
Бондаревский сквер
Боровецкий
Бородинская
Ботаническая
Брaтская
Брагинская
Брагинский 
Брагинский 
Браславская
Братский
Брест Литовская
Брестская
Брестский 
Брестский 
Брестский 
Брестский 
Брестский 
Брикета
Брилевская
Брилевский
Бровикова
Броневой
Тарашкевича
Будённого
Будславская
Будславский
Бумажкова
Бумажкова
Бумажкова
Бурдейного
Быховская
Вавилова
Вавилова
Валентия Ваньковича
Ванеева
Ванеева
Варвашени
Васильковая
Васнецова
Васнецова 
Васнецова 
Ватутина
Ватутина
Ватутина
Ваупшасова
Великоморская
Велозаводская
Велосипедный
Велосипедный 
Вербная
Вересковая
Верещагина
Верхний
Верхняя
Хоружей
Весенняя
Веснина
Веснинка
Веснинка
Взлётная
Турова
Вилейская
Вилковщина
Вилковщина
Вильямса
Вильямса
Вильямса 
Виноградная
Вирская
Витебская
Вишнёвый
Дубовки
Оловникова
Оловникова
Голубка
Сырокомли
Водозаборная
Водолажского
Водопроводный
Войсковый
Вокзальная
Вокзальный 
Волгоградская
Волгодонская
Волжская
Волжский
Володарского
Володько
Воложинская
Волоха
Волочаевская
Волочаевский
Волочаевский 
Воронянского
Воропаевская
Восточная
Восточный
Игнатовского
Всехсвятская
Встречная
Встречный
Вузовский
Выготского
Высокая
Высокий
Вязынская
Вязынский 
Вязынский 
Гагарина
Гало
Гамарника
Гастелло
Гая
Гвардейская
Гвардейский 
Гебелева
Антонова
Цитовича
Геологическая
Геологический
Геологический
Герасименко
Германовская
Германовская 
Германовский
Герцена
Гикало
Гинтовта
Глаголева
Глаголева 
Глаголева 
Глинище
Глубокская
Гоголевская
Головачёва
Голодеда
Голодеда
Голубева
Гольшанская
Гольшанский 
Гольшанский 
Горбатова
Горная
Горный
Горовца
Городецкая
Городской Вал
Гребельки
Грекова
Грекова 
Грекова 
Грибной
Грибоедова
Ширмы
Гризодубовой
Грицевца
Гродненская
Громова
Грушевская
Грушевский
Грушевский сквер
Гурского
Гуртьева
Гусовского
Дальний
Сердича
Даровская
Даровский
Даумана
Дача
Дачная
Дачный
Дачный
Двинская
Дворищи
Дежнёва
Декабристов
Демьяна Бедного
Денисовская
Денисовский
Деревенская
Детский 
Детский 
Детский 
Дзержинского
Дививелка
Димитрова
Дисенская
Днепровская
Днепровский
Добромысленский
Доватора
Докутович
Докучаева
Докшицкая
Долгиновский
Долгиновский
Долгиновский 
Долгиновский 
Долгиновский 
Долгобродская
Домашевская
Домашевский
Домбровская
Дорошевича
Достоевского
Достоевского
Дражня
Дрозда
Дрозды
Дружбы
Друйская
Дубовлянский
Дубравинская
Дубравинский
Дукорская
Дунина Марцинкевича
Глебова
Мирковского
Мирковского
Мирковского
Евфросиньи Полоцкой
Ежи Гедройца
Ельницкая
Ельских
Енисейский
Енисейский 
Енисейский 
Ермака
Ермака
Жасминовая
Ждановичская
Железнодорожная
Железнодорожный 
Железнодорожный 
Железнодорожный 
Железнодорожный 
Железнодорожный 
Железнодорожный 
Железнодорожный 
Жилуновича
Жлобинская
Жудро
Жукова
Жуковского
Жуковского
Жуковского 
Жуковского 
Заводская
Загородный 
Загородный 
Загородный 
Загородный 
Замковая
Западная
Запорожская
Запорожский 
Запорожский 
Запрудная
Заречанский
Заречная
Заречный
Заречный 
Заречный 
Заречный 
Заславская
Заслонова
Затишье
Захарова
Зацень
Звёздная
Зелёнолугская
Земледельческая
Земледельческий 
Земледельческий 
Землемерная 
Землемерная 
Землемерный
Зенитный
Зимний 
Зимний 
Зимняя
Змитрока Бядули
Зои Космодемьянской
Золотая
Золотая Горка
Золотой
Зубачёва
Зубачёва 
Зубачёва 
Зубачёва 
Зыбицкая
Науменко
Хруцкого
Шамякина
Ивановская
Ивенецкая
Буйницкого
Буйницкого
Домейко
Игнатенко
Игуменский
Иерусалимская
Извозная
Извозный 
Извозный 
Извозный 
Измайловская
Измайловский
Измайловский 
Измайловский 
Измайловский 
Измайловский 
Изумрудный
Илимская
Копиевича
Ильменская
Ильянская
Герасименко
Герасименко
Павлова
Чавеса
Айтматова
Индустриальная
Инженерная
Инженерный
Инструментальный
Интернациональная
Иодковская
Гошкевича
Жиновича
Ириновская
Ириновский
Иркутская
Иркутский
Искалиева
Кабушкина
Кабушкина
Казарменный
Казимировская
Казинца
Казинца
Калинина
Калинина
Калинина
Калининградский
Калиновского
Кальварийская
Кальварийский
Кальварийский
Камайская
Каменногорская
Канатный
Карбышева
Карвата
Либкнехта
Маркса
Каролинская
Каролинский
Карпова
Каганца
Каховская
Каховский
Качинская
Каштановая
Кедышко
Киевская
Киреева
Киреенко
Кирилла и Мефодия
Кирилла Туровского
Кирова
Киселёва
Клары Цеткин
Клецкая
Клецкий
Клубный
Клумова
Клумова
Ключевая
Кнорина
Княгининская
Княгининский
Кобринская
Ковалёва
Козлова
Козлова
Козыревская
Козьмо Демьяновский
Колесникова
Колесникова
Коллективный
Коллективный
Коллекторная
Колхозная
Кольцевая 
Кольцевая 
Кольцевой 
Кольцевой 
Кольцова
Кольцова
Кольцова 
Кольцова 
Кольцова 
Кольцова 
Колядная
Комаровский
Коммунистическая
Комсомольская
Крапивы
Буйло
Кооперативный
Копыльская
Копыльский
Копыльский 
Коржа
Корженевского
Корженевского
Корзюки
Корицкого
Короленко
Короля
Короткевича
Короткий 
Короткий 
Короткий 
Корш Саблина
Космонавтов
Котовского
Котовского
Крайняя
Крамского
Красивая
Красивый
Красная
Красная Слобода
Красноармейская
Краснодонская
Краснозвёздная
Краснозвёздный
Краснослободская
Красный
Кривичская
Кривичский 
Кривичский 
Кропоткина
Круглый
Крупецкая
Крупской
Крупцы
Крыжовская
Крыловича
Крыловича 
Крыловича 
Кубанская
Кузнечная
Минина
Минина
Чорного
Чорного
Куйбышева
Кулешова
Кулибина
Кульман
Кунцевщина
Куприянова
Курганная
Гвишиани
Кутузова
Кэчевский
Лазарева
Лазо
Ландера
Лапоровичская
Александровской
Левкова
Кижеватова
Ленина
Ленинградская
Леонида Беды
Лепельская
Лермонтова
Леси Украинки
Лесопарковая
Летний 
Летний 
Летняя
Лещинского
Либаво Роменская
Ливенцева
Лидская
Чайкиной
Карастояновой
Линейный
Липковская
Липковский
Липковский
Липовая
Литературная
Литературный
Лобанка
Логойский
Ложинская
Ломоносова
Лошица 
Лошица 
Лошицкая
Лошицкий
Лошицкий
Луговая
Луговой
Лукьяновича
Лучайская
Сапеги
Толстого
Любанская
Любимова
Лютеранский
Люцинская
Ляховский сквер
Магнитная
Магнитный
Мазурова
Майкова
Макаёнка
Макаёнка
Макаренко
Богдановича
Богдановича
Горецкого
Горького
Танка
Малакович
Малая
Малинина
Малое Стиклево
Маломедвежинский
Малосеребрянская
Малофеевская
Малый
Малый 
Малый Тростенец
Малявки
Маневича
Марата
Маргелова
Марусинская
Марусинский 
Марусинский 
Лосика
Марьевская
Масюковщина
Масюковщина
Масюковщина
Матвеевская
Матвеевский
Матросова
Матусевича
Мачульского
Машерова
Машинистов
Машинистов
Машиностроителей
Маяковского
Маяковского
Медвежино
Медвежино
Мележа
Мелиоративная
Мелиоративный
Мельникайте
Мельничный
Менделеева
Менделеева 
Мержинского
Металлистов
Миколаевская
Микулича
Минская
Мира
Мирная
Мирошниченко
Пташука
Михайлова
Михайловский
Михалoвская
Михалово
Михалово 
Михалово 
Михалово 
Михалово 
Лынькова
Чарота
Мичурина
Могилёвская
Могилёвское
Можайского
Можайского 
Можайского 
Можайского 
Мозырская
Мозырский
Мозырский
Молодечненская
Молочный
Монтажников
Монтажников 
Монтажников 
Монтажников 
Монтажников 
Москвина
Москвина
Московская
Музыкальный
Мулявина
Мядельская
Мясникова
Мястровская
Набережная
Навуковая
Нагорная
Нагорный
Надеждинская
Наклонная
Наклонный 
Наклонный 
Налибокская
Наполеона Орды
Народная
Нарочанская
Натуралистов
Нахимова
Нахимова
Невский
Неждановой
Неждановой
Независимости
Независимости
Независимости
Некрасова
Некрасова
Некрашевича
Нёманская
Немига
Несвижская
Нестерова
Нестерова 
Нестерова 
Нефтяная
Нефтяной
Никитина
Никитина 
Никитина 
Никифорова
Новаторская
Новаторский
Новаторский
Новая
Новгородская
Новгородский
Новинковская
Нововиленская
Нововиленский
Новосельская
Новоуфимская
Ногина
Обойная
Обойный
Обуховская
Обуховский 
Обуховский 
Обуховский 
Огарёва
Огинского
Огородницкая
Одесская
Одесский
Одинцова
Одинцова
Одоевского
Одоевского
Озёрная
Озёрный
Озерцовский
Окрестина
Окрестина 
Окружной
Октябрьская
Октябрьская
Кошевого
Олешева
Ольховая
Ольшанская
Ольшевский
Ольшевского
Ольшевского
Омельянюка
Оранжерейная
Орджоникидзе
Орловская
Орловский 
Орловский 
Орловский 
Орловский 
Орловский 
Орловский 
Оршанская
Оршанский
Освейская
Освобождения
Осенний 
Осенний 
Осенняя
Осипенко
Осиповичская
Основателей
Острожских
Острошицкая
Охотская
Охотский
Ошмянский
Шпилевского
Медёлки
Павлова
Павлова
Павловского
Труса
Панфилова
Папанина
Папернянская
Папернянский
Парашютная
Парижской Коммуны
Парковая
Парниковая
Парниковый 
Паровозный
Партизанская
Партизанский
Пархоменко
Пензенская
Первомайская
Передовая
Передовой
Переходная
Переходной 
Переходной 
Пермская
Пермский
Песочная 
Глебки
Мстиславца
Румянцева
Сергиевича
Петровщина
Петровщина
Петруся Бровки
Пилотская
Пильницкая
Панченко
Пинская
Пионерская
Пирогова
Пирогова 
Пирогова 
Письменников
Планёрная
Платонова
Плеханова
Победителей
Победы
Победы
Подгорная
Подлесная
Подольская
Подольский 
Подольский 
Подольский 
Подольский 
Подшипниковый
Подшипниковый 
Подшипниковый 
Пожарная
Пожарского
Покровская
Покровский
Полевая
Полевой
Полесская
Ползунова
Полиграфическая
Полиграфический 
Полиграфический 
Полиграфический 
Полиграфический 
Полоцкая
Полоцкий
Полтавская
Полярная
Пономарёва
Пономаренко
Попова
Поселковая 
Поселковая 
Поселковый
Поселковый 
Поселковый 
Поставская
Почтовая
Пржевальского
Пржевальского 
Привабная
Приветливая
Привокзальная
Пригородная
Придорожная
Прилукская
Прилукский 
Приозёрная
Притыцкого
Программистов
Прогрессивная
Проездной
Промышленная
Промышленный
Профсоюзная
Профсоюзный
Профсоюзный
Прушинских
Прямая
Псковская
Пугачёвская
Пулихова
Путейская
Путейский
Путепроводный 
Путепроводный 
Путепроводный 
Путепроводный 
Путепроводный 
Путепроводный 
Путепроводный 
Путепроводный 
Путилова
Пуховичская
Пушкина
Рaтомский 
Рaтомский 
Рабкоровская
Рабочий
Радашковская
Радашковский
Радиальная
Радиаторная 
Радиаторный 
Радиаторный 
Радиаторный 
Радиаторный 
Радиаторный 
Радищева
Радужная
Разинская
Разинский
Раковская
Рассонская
Ратомская
Раубичская
Раубичский
Рафиева
Революционная
Репина
Республикaнская
Речная
Ржавецкая
Ржавецкий
Рижская
Рижский 
Рижский 
Рогачёвская
Рогачёвский
Рогачёвский
Родниковая
Розы Люксембург
Рокоссовского
Романовская Слобода
Ромашкина
Роменская
Ротмистрова
Русановича
Руссиянова
Рыбалко
Рылеева
Рябинницкая
Рябиновая
Садовая
Самарский
Сапёров
Свердлова
Светлая
Свирский 
Свирский 
Свирский 
Свирский 
Свислочская
Свислочский
Свободы
Связистов
Севастопольская
Севастопольский
Северный
Седова
Седых
Селицкого
Семашко
Семёнова
Сёмковская
Сёмковский
Сенницкая
Сенницкий
Серафимовича
Есенина
Серебрянская
Серебрянский
Серова
Сеченова
Силикатный 
Силикатный 
Сиреневая
Скалинская
Мицкевича
Жукова
Притыцкого
Содружество
Скрипникова
Скрыганова
Славинского
Славный
Славянская
Слепнянский 
Слесарная
Слободская
Слободской
Слободской
Слонимская
Слуцкой
Смиловичский
Смирнова
Смоленская
Смолячкова
Сморговский
Сморговский
Сморговский 
Сморговский 
Сморговский 
Сморговский 
Сморговский 
Снежный
Собинова
Собинова
Собинова 
Советская
Соколянский
Солнечная
Солнечный
Соловьиный
Соломенная
Солтыса
Солтыса
Сосновая
Сосновый Бор
Ковалевской
Социалистическая
Спортивный
Стaйковская
Стадионная
Монюшко
Станиславского
Станционный
Стариновская
Старобинская
Старовиленская
Старовиленский
Старостинская слобода
Старотроицкая
Стасова
Стасова 
Стасова 
Стасова 
Стахановская
Стахановский
Стебенёва
Стебенёва
Стекольный
Степной
Степянская
Степянский
Стефании Станюты
Стефановская
Столбцовская
Столетова
Столетова
Столетова
Столичная
Столичный
Столичный
Сторожовская
Стрелковая
Строителей
Струговская
Студенческая
Студенческий
Суворова
Судмалиса
Суражская
Сурганова
Сурикова
Сурикова
Сурикова
Сухаревская
Сухая
Таёжная
Талаша
Талаша 
Таллинская
Таллинский
Танковая
Тарханова
Татарская
Татарский
Ташкентская
Ташкентский
Твёрдый 
Телеграфный
Тенистая
Тепличная
Тепличный
Тикоцкого
Тимирязева
Тимошенко
Тимошенко 
Тимошенко 
Тиражная
Тиражный
Тиражный 
Тиражный 
Тихая
Тихий
Тобольский
Товарищеский
Толбухина
Томская
Томский
Топографическая
Топографический 
Топографический 
Торговый
Тракторостроителей
Троицкая
Тростенецкая
Трубный 
Трубный 
Трубный 
Трудовая
Тульская
Тульский 
Тульский 
Тупиковая
Туполева
Тургенева
Тухачевского
Тучинский
Тышкевичей
Тяпинского
Уборевича
Украинская
Ульяновская
Уманская
Уманская 
Университетский
Уральская
Уральский
Урожайная
Уручская
Усадебная
Усадебный
Утульная
Уфимская
Ученический
Ушакова
Ушакова
Ушачская
Фабрициуса
Фабрициуса
Фабричная
Фабричный 
Фабричный 
Фанипольская
Фанипольский
Федотова
Федотова
Рущица
Физкультурная
Физкультурный
Физкультурный 
Филатова
Филимонова
Фогеля
Фольварковый
Фомина
Фомина 
Фомина 
Богушевича
Скорины
Скорины 
Фроликова
Фрунзе
Фучика
Хабаровская
Халтурина
Халтурина 
Халтурина 
Харьковская
Харьковский 
Хмаринская
Хмелевского
Хмельницкого
Холмогорская
Холмогорский 
Холмогорский 
Холмогорский 
Холмогорский 
Центральная
Центральная
Цнянская
Чайковского
Чайковского
Чайковского
Чайлытко
Чапаева
Чеботарёва
Чекалина
Чекалина
Челюскинцев
Челюскинцев 
Червякова
Червякова
Черниговская
Черниговский
Чернышевского
Чернышевского 
Черняховского
Чехова
Чигладзе
Чижевских
Чижевских
Чижевских
Чижовская
Чичерина
Чичурина
Чкалова
Чюрлёниса
Шабаны
Шабловского
Шаранговича
Шатило
Шатько
Шафарнянская
Шевченко
Шейпичи
Шестая Линия 
Широкая
Шишкина
Школьная
Школьный
Шорная
Шоссейная
Шугаева
Шумилинская
Щедрина
Щедрина 
Щедрина 
Щедрина 
Щепкина
Щербакова
Щербакова
Щётовка
Щорса
Щорса 
Щорса 
Щорса 
Щорса 
Щукина
Экскаваторная
Экскаваторный
Энгельса
Энергетический
Юбилейная
Юбилейная
Юбилейный
Южная
Южный
Южный 
Юношеская
Юношеский 
Юношеский 
Семеняко
Юрово Завальная
Яблоневая
Дроздовича
Якуба Коласа
Якубова
Якубовского
Борщевского
Райниса
Чечота
Брыля
Купалы
Лучины
Мавра
Янковского
Янтарная
Янтарный
Ярковская
Ярошевичская
Ярошевичский
Ясная
Автозаводская
Академия наук
Борисовский тракт
Восток
Институт культуры
Каменная Горка
Кунцевщина
Купаловская
Могилевская
Молодежная
Московская
Немига
Октябрьская
Парк Челюскинцев
Партизанская
Первомайская
Площадь Ленина
Площадь Победы
Площадь Якуба Коласа
Пролетарская
Пушкинская
Спортивная
Тракторный завод
Уручье
Фрунзенская`