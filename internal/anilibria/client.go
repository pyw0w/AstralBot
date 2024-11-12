package anilibria

// GetTitle получает информацию о тайтле по id или коду
// Параметры:
// - id: идентификатор тайтла
// - code: код тайтла
// Пример использования:
// params := map[string]string{"id": "123"}
// title, err := GetTitle(params)
// Возвращаемые объекты:
// - names: объект с названиями тайтла (ru, en, alternative)
// - posters: объект с постерами (small, medium, original)
// - episodes: объект с информацией о сериях (string, first, last)
// - status: объект со статусом тайтла (string, code)
// - type: объект с типом тайтла (full_string, string, episodes, length, code)
// - team: объект с командой (voice, translator, editing, decor, timing)
// - season: объект с сезоном (year, week_day, string, code)
// - franchises: массив объектов с информацией о франшизах
// - blocked: объект с информацией о блокировке (blocked, bakanim)
// - player: объект с информацией о плеере (alternative_player, host, list, rutube, episodes)
// - torrents: объект с информацией о торрент файлах (episodes, list)
func GetTitle(params map[string]string) (string, error) {
	return fetchData("https://api.anilibria.tv/v3/title", params)
}

// GetTitleList получает информацию о нескольких тайтлах сразу
// Параметры:
// - id_list: список идентификаторов тайтлов через запятую
// Пример использования:
// params := map[string]string{"id_list": "123,456,789"}
// titles, err := GetTitleList(params)
// Возвращаемые объекты:
// - см. GetTitle
func GetTitleList(params map[string]string) (string, error) {
	return fetchData("https://api.anilibria.tv/v3/title/list", params)
}

// GetTitleUpdates получает список тайтлов, отсортированных по времени добавления нового релиза
// Параметры:
// - limit: количество тайтлов
// Пример использования:
// params := map[string]string{"limit": "10"}
// updates, err := GetTitleUpdates(params)
// Возвращаемые объекты:
// - см. GetTitle
func GetTitleUpdates(params map[string]string) (string, error) {
	return fetchData("https://api.anilibria.tv/v3/title/updates", params)
}

// GetTitleChanges получает список тайтлов, отсортированных по времени изменения
// Параметры:
// - limit: количество тайтлов
// Пример использования:
// params := map[string]string{"limit": "10"}
// changes, err := GetTitleChanges(params)
// Возвращаемые объекты:
// - см. GetTitle
func GetTitleChanges(params map[string]string) (string, error) {
	return fetchData("https://api.anilibria.tv/v3/title/changes", params)
}

// GetTitleSchedule получает расписание выхода тайтлов, отсортированное по дням недели
// Параметры:
// - day: день недели (например, "monday")
// Пример использования:
// params := map[string]string{"day": "monday"}
// schedule, err := GetTitleSchedule(params)
// Возвращаемые объекты:
// - см. GetTitle
func GetTitleSchedule(params map[string]string) (string, error) {
	return fetchData("https://api.anilibria.tv/v3/title/schedule", params)
}

// GetTitleRandom возвращает случайный тайтл из базы
// Параметры отсутствуют
// Пример использования:
// randomTitle, err := GetTitleRandom()
// Возвращаемые объекты:
// - см. GetTitle
func GetTitleRandom() (string, error) {
	return fetchData("https://api.anilibria.tv/v3/title/random", nil)
}

// GetTitleSearch возвращает список найденных по фильтрам тайтлов
// Параметры:
// - search: строка поиска
// Пример использования:
// params := map[string]string{"search": "Naruto"}
// searchResults, err := GetTitleSearch(params)
// Возвращаемые объекты:
// - см. GetTitle
func GetTitleSearch(params map[string]string) (string, error) {
	return fetchData("https://api.anilibria.tv/v3/title/search", params)
}

// GetTitleSearchAdvanced выполняет поиск информации по продвинутым фильтрам с поддержкой сортировки
// Параметры:
// - search: строка поиска
// - sort: поле для сортировки
// Пример использования:
// params := map[string]string{"search": "Naruto", "sort": "rating"}
// advancedSearchResults, err := GetTitleSearchAdvanced(params)
// Возвращаемые объекты:
// - см. GetTitle
func GetTitleSearchAdvanced(params map[string]string) (string, error) {
	return fetchData("https://api.anilibria.tv/v3/title/search/advanced", params)
}

// GetTitleFranchises получает информацию о франшизе по ID тайтла
// Параметры:
// - id: идентификатор тайтла
// Пример использования:
// params := map[string]string{"id": "123"}
// franchises, err := GetTitleFranchises(params)
// Возвращаемые объекты:
// - franchises: массив объектов с информацией о франшизах
func GetTitleFranchises(params map[string]string) (string, error) {
	return fetchData("https://api.anilibria.tv/v3/title/franchises", params)
}

// GetYoutube получает информацию о вышедших роликах на YouTube каналах в хронологическом порядке
// Параметры отсутствуют
// Пример использования:
// youtubeInfo, err := GetYoutube()
// Возвращаемые объекты:
// - информация о роликах на YouTube
func GetYoutube() (string, error) {
	return fetchData("https://api.anilibria.tv/v3/youtube", nil)
}

// GetFeed получает список обновлений тайтлов и роликов на YouTube каналах в хронологическом порядке
// Параметры отсутствуют
// Пример использования:
// feed, err := GetFeed()
// Возвращаемые объекты:
// - информация об обновлениях тайтлов и роликов на YouTube
func GetFeed() (string, error) {
	return fetchData("https://api.anilibria.tv/v3/feed", nil)
}

// GetYears возвращает список годов выхода доступных тайтлов по возрастанию
// Параметры отсутствуют
// Пример использования:
// years, err := GetYears()
// Возвращаемые объекты:
// - список годов выхода тайтлов
func GetYears() (string, error) {
	return fetchData("https://api.anilibria.tv/v3/years", nil)
}

// GetGenres возвращает список всех жанров по алфавиту
// Параметры отсутствуют
// Пример использования:
// genres, err := GetGenres()
// Возвращаемые объекты:
// - список жанров
func GetGenres() (string, error) {
	return fetchData("https://api.anilibria.tv/v3/genres", nil)
}

// GetTeam возвращает список участников команды, когда-либо существовавших на проекте
// Параметры отсутствуют
// Пример использования:
// team, err := GetTeam()
// Возвращаемые объекты:
// - список участников команды
func GetTeam() (string, error) {
	return fetchData("https://api.anilibria.tv/v3/team", nil)
}

// GetTorrentSeedStats возвращает список пользователей и их статистику на трекере
// Параметры отсутствуют
// Пример использования:
// seedStats, err := GetTorrentSeedStats()
// Возвращаемые объекты:
// - статистика пользователей на трекере
func GetTorrentSeedStats() (string, error) {
	return fetchData("https://api.anilibria.tv/v3/torrent/seed_stats", nil)
}

// GetTorrentRSS возвращает список обновлений на сайте в одном из форматов RSS ленты
// Параметры:
// - format: формат RSS ленты (например, "xml")
// Пример использования:
// params := map[string]string{"format": "xml"}
// rss, err := GetTorrentRSS(params)
// Возвращаемые объекты:
// - информация об обновлениях в формате RSS
func GetTorrentRSS(params map[string]string) (string, error) {
	return fetchData("https://api.anilibria.tv/v3/torrent/rss", params)
}

// GetFranchiseList возвращает список всех франшиз
// Параметры отсутствуют
// Пример использования:
// franchiseList, err := GetFranchiseList()
// Возвращаемые объекты:
// - список франшиз
func GetFranchiseList() (string, error) {
	return fetchData("https://api.anilibria.tv/v3/franchise/list", nil)
}
