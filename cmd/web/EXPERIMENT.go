package web

/*
// Создается функция-обработчик "home", которая записывает байтовый слайс, содержащий
// текст "Привет из Snippetbox" как тело ответа.
func home(w http.ResponseWriter, r *http.Request) {
	// Параметр http.ResponseWriter предоставляет методы для объединения
	//HTTP ответа и возвращение его пользователю, а второй параметр
	//*http.Request является указателем на структуру, которая содержит информацию
	// о текущем запросе (вроде HTTP-методов POST, GET, DELETE… и URL текущего запроса)
	w.Write([]byte("Привет из Snippetbox"))
}

func main() {
	// Используется функция http.NewServeMux() для инициализации нового рутера, затем
	// функцию "home" регистрируется как обработчик для URL-шаблона "/".
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Используется функция http.ListenAndServe() для запуска нового веб-сервера.
	// Мы передаем два параметра: TCP-адрес сети для прослушивания (в данном случае это "localhost:4000")
	// и созданный рутер. Если вызов http.ListenAndServe() возвращает ошибку
	// мы используем функцию log.Fatal() для логирования ошибок. Обратите внимание
	// что любая ошибка, возвращаемая от http.ListenAndServe(), всегда non-nil.
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}*/
// Обработчик главной страницы.
/*
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" { // proverka addresa
		http.NotFound(w, r) // err 404
		return
	}
	w.Write([]byte("Привет из Snippetbox"))
}

// Обработчик для отображения содержимого заметки.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Отображение выбранной заметки с ID %d...", id)
	w.Write([]byte(string(id)))
}
*/
/*
// Обработчик для создания новой заметки.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		w.Header().Set("Allow", http.MethodPost) // Вы должны сперва вызвать метод w.Header().Set() и уже потом остальные методы w.WriteHeader() или w.Write() иначе ваши изменения не будут учтены.

		//w.WriteHeader(405)//Vызвать метод w.WriteHeader() в обработчике можно только один раз, и после возвращения кода состояния HTTP, изменить его нельзя. Если попытаться вызвать w.WriteHeader() во второй раз, Go выдаст сообщение об ошибке;
		//w.Write([]byte("GET-Метод запрещен!"))
		//return*

		http.Error(w, "Метод запрещен!", 405) // alternativa
		return

	}
	w.Write([]byte("Форма для создания новой заметки..."))
}
*/

/*
func main() {
	// Регистрируем два новых обработчика и соответствующие URL-шаблоны в
	// маршрутизаторе servemux
	mux := http.NewServeMux()
	// Фиксированные пути не заканчиваются косой чертой, тогда как многоуровневые пути заканчиваются косой чертой.
	mux.HandleFunc("/", home) // роутинг "/" является примером многоуровневого пути
	// Поскольку DefaultServeMux является глобальной переменной,
	// любой пакет может получить к ней доступ и зарегистрировать маршрут — включая любые сторонние пакеты, которые использует ваше приложение.
	// http.HandleFunc("/snippet", showSnippet) //default HandleMux
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	// У многоуровневых путей в конце может быть какой-то вспомогательный символ. К примеру, "/**" или "/static/**"

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000/snippet")
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000/snippet/create")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
*/
