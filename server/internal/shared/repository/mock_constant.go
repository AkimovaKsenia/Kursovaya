package repository

import "kino/internal/shared/entities"

var (
	mockRoles = []entities.Role{
		{
			Name: "worker",
		},
		{
			Name: "admin",
		},
	}

	mockUsers = []entities.User{
		{
			Name:     "Данек",
			Surname:  "Гевинов",
			Email:    "worker@yandex.ru",
			Password: "worker",
			RoleID:   1,
		},
		{
			Name:     "Ксесша",
			Surname:  "Акимова",
			Email:    "admin@yandex.ru",
			Password: "admin",
			RoleID:   2,
		},
	}

	mockFilmStudios = []entities.FilmStudio{
		{Name: "Walt Disney Pictures"},
		{Name: "Universal Pictures"},
		{Name: "Paramount Pictures"},
		{Name: "Sony Pictures"},
		{Name: "Warner Bros. Pictures"},
		{Name: "Мосфильм"},
		{Name: "СТВ"},
		{Name: "Central Partnership"},
		{Name: "20th Century Studios"},
		{Name: "Газпром-Медиа"},
		{Name: "Творческое объединение Мистерия"},
		{Name: "Централ Партнершип"},
	}

	mockGenres = []entities.Genre{
		{Name: "Мультфильм"},
		{Name: "Фэнтези"},
		{Name: "Приключения"},
		{Name: "Комедия"},
		{Name: "Ужасы"},
		{Name: "Триллер"},
		{Name: "Боевик"},
		{Name: "Драма"},
		{Name: "Семейный"},
		{Name: "Детектив"},
		{Name: "Военный"},
		{Name: "История"},
		{Name: "Спорт"},
	}

	mockDirectors = []entities.Director{
		{FIO: "Джоэль Кроуфорд"},
		{FIO: "Джаред Буш"},
		{FIO: "Джон Красински"},
		{FIO: "Ли Кронайдер"},
		{FIO: "Мэтт Беттинелли-Олпин"},
		{FIO: "Тайлер Джиллетт"},
		{FIO: "Дэвид Яровески"},
		{FIO: "Алексей Нужный"},
		{FIO: "Сергей Коротаев"},
		{FIO: "Петр Зеленов"},
		{FIO: "Владимир Пресняков"},
		{FIO: "Олег Пресняков"},
		{FIO: "Николай Лебедев"},
		{FIO: "Яна Кузьмина"},
		{FIO: "Илья Учитель"},
		{FIO: "Лоуренс Фаулер"},
	}

	mockOperators = []entities.Operator{
		{FIO: "Майкл Бургер"},
		{FIO: "Джон Шварцман"},
		{FIO: "Дэн Лаустсен"},
		{FIO: "Брэндон Трост"},
		{FIO: "Аарон Мортон"},
		{FIO: "Илья Демин"},
		{FIO: "Андрей Найденов"},
		{FIO: "Игорь Гринякин"},
		{FIO: "Андрей Найденов"},
		{FIO: "Дмитрий Трифонов"},
		{FIO: "Сергей Мачильский"},
		{FIO: "Павел Медведев"},
		{FIO: "Нил Парсонс"},
	}

	mockFilms = []entities.CreateFilm{
		{
			Name:          "В списках не значился",
			Description:   "21 июня 1941 года. Полный надежд на светлое будущее молодой лейтенант Коля Плужников (Владислав Миллер) прибывает в Брест, чтобы начать службу. Однако судьба готовит ему испытание: солдата не успевают зачислить в личный состав военнослужащих, а в четыре утра раздаются артиллерийские разрывы — начинается война. Так Коля Плужников принимает первый в своей жизни бой продолжительностью в десять месяцев. ",
			Photo:         "https://s1ru1.kinoplan24.ru/670/0406060506d9b313a274d382/21460.jpg?mode=fit&width=1024&height=1024",
			CastList:      []string{"Владислав Миллер", "Алена Морилова", "Владимир Машков", "Павел Чернышев", "Яна Сексте"},
			FilmStudioID:  10,
			DurationInMin: 120,
			DirectorIDs:   []int{9},
			OperatorIDs:   []int{8},
			GenreIDs:      []int{8, 11, 12},
		},
		{
			Name:          "Финал",
			Description:   "12 августа 2012 года. Выставочный центр «Эрлс-Корт» в Лондоне наполняется болельщиками. Миллионы зрителей по всему миру ждут начала трансляции финальной игры по волейболу среди мужских команд, в которой определится победитель ХХХ Олимпийских игр. Сборная России, никогда не бравшая золото. И сборная Бразилии, обласканная контрактами, победами и преданными фанатами. Многим кажется, что результат игры предопределён. Но ещё мало кто знает, какое место занимает этот финал в жизни каждого из российских спортсменов.",
			Photo:         "https://s1ru1.kinoplan24.ru/671/0406060506d71e4e26a4ad9d/21435.jpg?mode=fit&width=1024&height=1024",
			CastList:      []string{"Олег Куликович", "Сергей Гармаш", "Ольга Кузьмина"},
			FilmStudioID:  11,
			DurationInMin: 96,
			DirectorIDs:   []int{10, 11, 12},
			OperatorIDs:   []int{9, 10},
			GenreIDs:      []int{8, 13},
		},
		{
			Name:          "Кракен",
			Description:   "Российский ракетный подводный крейсер специального назначения бесследно исчезает во время секретной миссии в Гренландском море. На его поиски отправляется экипаж под командованием Виктора Воронина, чей старший брат Александр Воронин командует пропавшей субмариной.",
			Photo:         "https://s1ru1.kinoplan24.ru/670/0406060506d196eea7602d47/21479.jpg?mode=fit&width=1024&height=1024",
			CastList:      []string{"Александр Петров", "Алексей Гуськов", "Сергей Гармаш", "Виктор Добронравов", "Диана Пожарская"},
			FilmStudioID:  12,
			DurationInMin: 134,
			DirectorIDs:   []int{13},
			OperatorIDs:   []int{11},
			GenreIDs:      []int{7, 8},
		},
		{
			Name:          "Мальчик-Дельфин 2",
			Description:   "Героические подвиги Мальчика-Дельфина сделали его легендой среди морских обитателей. Ему предстоит не только спасти своих друзей, но и раскрыть тайны своего прошлого.",
			Photo:         "https://s1ru1.kinoplan24.ru/664/0406060506db0a9c2f661ba6/23200.jpg?mode=fit&width=1024&height=1024",
			CastList:      []string{"Никита Кологривый", "Полина Авдеенко", "Юлия Рудина", "Юлия Зоркина", "Александр Васильев"},
			FilmStudioID:  7,
			DurationInMin: 95,
			DirectorIDs:   []int{14},
			GenreIDs:      []int{1},
		},
		{
			Name:          "Батя 2. Дед",
			Description:   "Макс и Ирина находятся на грани развода. Их сын Дима сбежал в деревню к Бате. В пути Макс вспоминает свое детство с Дедом.",
			Photo:         "https://s2ru1.kinoplan24.ru/1373/0406060506bd5732700c4a91/22580.jpg?mode=fit&width=1024&height=1024",
			CastList:      []string{"Владимир Вдовиченков", "Стас Старовойтов", "Евгений Цыганов", "Надежда Михалкова", "Андрей Андреев"},
			FilmStudioID:  12,
			DurationInMin: 89,
			DirectorIDs:   []int{15},
			OperatorIDs:   []int{12},
			GenreIDs:      []int{4},
		},
		{
			Name:          "Кукла. Реинкарнация",
			Description:   "После смерти сына эксперт по робототехнике создает аниматронную куклу, которая невероятно похожа на его умершего ребенка. Но у Робина есть свои собственные намерения...",
			Photo:         "https://s2ru1.kinoplan24.ru/666/0406060506d11eef423e3250/22283.jpg?mode=fit&width=1024&height=1024",
			CastList:      []string{"Итан Тейлор", "Михаэла Лонгден", "Леона Кларк", "Саймон Дэвис", "Виктор Меллорс"},
			FilmStudioID:  4,
			DurationInMin: 91,
			DirectorIDs:   []int{16},
			OperatorIDs:   []int{13},
			GenreIDs:      []int{5, 6},
		},
	}

	mockCinemaConditions = []entities.CinemaCondition{
		{Name: "Работает в обычном режиме"},
		{Name: "Открыт с ограничениями"},
		{Name: "Закрыт на санитарный день"},
		{Name: "Закрыт на техническое обслуживание"},
		{Name: "Реконструкция здания"},
		{Name: "Временное приостановление работы"},
		{Name: "Открытие после ремонта"},
		{Name: "Работает только онлайн-продажи"},
		{Name: "Ограниченный режим работы"},
		{Name: "Закрыт по решению администрации"},
	}

	mockCinemaCategories = []entities.CinemaCategory{
		{Name: "Стандартный"},
		{Name: "Артхаусный"},
		{Name: "IMAX"},
		{Name: "Премиум-класса"},
		{Name: "Детский"},
		{Name: "Летний открытый"},
		{Name: "Автокинотеатр"},
		{Name: "Обзорный"},
	}

	mockCinemaHallTypes = []entities.CinemaHallType{
		{Name: "Стандартный зал"},
		{Name: "VIP-зал"},
		{Name: "IMAX"},
		{Name: "3D"},
		{Name: "Dolby Atmos"},
	}

	mockCinemas = []entities.CreateCinema{
		{
			Name:        "КИНО ОККО",
			Description: "КИНО ОККО в «Афимолл Сити» – это современный кинотеатр крупнейшей сети «Синема парк и Формула кино».\n\nКинотеатр для тех, кто ценит качество и постоянство.\n\nМедиацентр главных событий киноиндустрии и бизнес-среды Москвы.\n\nЗалы оснащены инновационным оборудованием, которые обеспечивают захватывающий эффект полного погружения.\n\nПремиум зал Screen Max Sapphire – сверхмощная графика, объемный звук, привилегированное обслуживание с собственным лаундж-баром.",
			Photo:       "https://afimall.ru/netcat_files/390_1939.jpg",
			Address:     "Москва, Пресненская наб., 2, ТЦ «Афимолл-сити», 5-й этаж",
			Email:       "gewinn@cherepok.ksesha",
			Phone:       "+7(495)419-21-07",
			ConditionID: 1,
			CategoryID:  3,
		},
		{
			Name:        "Киномакс Жулебино",
			Description: "Кинотеатр в Москве «Киномакс-Жулебино» – это 6 стильных комфортабельных залов на 762 места, оснащенные современным оборудованием и удобными креслами, среди которых зал со звуком Dolby Atmos и  VIP-зал.",
			Photo:       "https://avatars.mds.yandex.net/get-altay/2057543/2a0000016d24c41222c9aaef5ffea0c164cf/XXXL",
			Address:     "г. Москва, ул. Генерала Кузнецова, 22, ТРЦ «Миля»",
			Email:       "zhulebino@kinomax.ru",
			Phone:       "+7(952)812-02-02",
			ConditionID: 1,
			CategoryID:  1,
		},
		{
			Name:        "Кинозал ГУМа",
			Description: "Кинозал ГУМа — это камерный театр из трех залов, где показывают кино. Архитектурный проект был выполнен по специальному заказу ГУМа. Классический «театральный декор» отвечает самым высоким акустическим стандартам привычных «черных коробок».",
			Photo:       "https://gum.ru/local/templates/gum_main/images/kinozal/kino-new1-min.jpg",
			Address:     "г. Москва, ул. Генерала Кузнецова, 22, ТРЦ «Миля»",
			Email:       "info@gum.ru",
			Phone:       "+7 (495) 788-43-43",
			ConditionID: 1,
			CategoryID:  4,
		},
		{
			Name:        "Синема парк Deluxe",
			Description: "Кинотеатр «Синема Парк DELUXE» расположен вблизи станции метро Войковская, в торговом центре «Метрополис».\n\nОсобенность дизайна этого кинотеатра заключается в том, что интерьер стилизован под зону курортного отдыха — на стенах висят большие фотографии с видом на Майами. Современный стиль хай-тек позволит вам расслабиться и отвлечься от суеты повседневной жизни.",
			Photo:       "https://avatars.mds.yandex.net/get-altay/5583647/2a000001804243d9ea7d2433d85ca404c46f/XXXL",
			Address:     "ул. Мультяшная, 5",
			Email:       "metropolis@cinema.ru",
			Phone:       "+7 (495) 197-77-11",
			ConditionID: 1,
			CategoryID:  4,
		},
	}

	mockCinemaHalls = []entities.CinemaHall{
		{
			Name:     "Зал 1",
			Capacity: 100,
			TypeID:   1,
			CinemaID: 1,
		},
		{
			Name:     "Зал 2",
			Capacity: 100,
			TypeID:   1,
			CinemaID: 1,
		},
		{
			Name:     "Зал 3",
			Capacity: 100,
			TypeID:   1,
			CinemaID: 1,
		},
		{
			Name:     "Зал 4",
			Capacity: 90,
			TypeID:   3,
			CinemaID: 1,
		},
		{
			Name:     "Зал 5",
			Capacity: 90,
			TypeID:   3,
			CinemaID: 1,
		},
		{
			Name:     "Зал 6",
			Capacity: 80,
			TypeID:   4,
			CinemaID: 1,
		},
		{
			Name:     "Зал 7",
			Capacity: 80,
			TypeID:   4,
			CinemaID: 1,
		},
		{
			Name:     "Зал 1",
			Capacity: 110,
			TypeID:   1,
			CinemaID: 2,
		},
		{
			Name:     "Зал 2",
			Capacity: 110,
			TypeID:   1,
			CinemaID: 2,
		},
		{
			Name:     "Зал 3",
			Capacity: 120,
			TypeID:   1,
			CinemaID: 2,
		},
		{
			Name:     "Зал 4",
			Capacity: 120,
			TypeID:   1,
			CinemaID: 2,
		},
		{
			Name:     "Зал 5",
			Capacity: 70,
			TypeID:   2,
			CinemaID: 2,
		},
		{
			Name:     "Зал 6",
			Capacity: 70,
			TypeID:   2,
			CinemaID: 2,
		},
		{
			Name:     "Большой зал",
			Capacity: 70,
			TypeID:   2,
			CinemaID: 3,
		},
		{
			Name:     "Детский зал",
			Capacity: 20,
			TypeID:   1,
			CinemaID: 3,
		},
		{
			Name:     "Специальный зал",
			Capacity: 16,
			TypeID:   2,
			CinemaID: 3,
		},
		{
			Name:     "зал Киносалон",
			Capacity: 25,
			TypeID:   2,
			CinemaID: 3,
		},
		{
			Name:     "Зал 1",
			Capacity: 100,
			TypeID:   4,
			CinemaID: 4,
		},
		{
			Name:     "Зал 2",
			Capacity: 100,
			TypeID:   4,
			CinemaID: 4,
		},
		{
			Name:     "Зал 3",
			Capacity: 100,
			TypeID:   4,
			CinemaID: 4,
		},
		{
			Name:     "Зал 4",
			Capacity: 100,
			TypeID:   4,
			CinemaID: 4,
		},
		{
			Name:     "Зал 5",
			Capacity: 32,
			TypeID:   2,
			CinemaID: 4,
		},
		{
			Name:     "Зал 6",
			Capacity: 32,
			TypeID:   2,
			CinemaID: 4,
		},
	}
)
