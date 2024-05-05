# Protobuf simple example for introduce in technology.

Простой пример создания `proto` файла и генерации в go код. Пример условно разделен на две части.

* В одном исполняемом файле (add_person.go) на go создается программа которая читает файл переданный в качестве аргумента и парсит его согласно структуре данных полченной из proto описания. Так же в коде реализован простой пример интерактивного взаимодействия с пользователем через прослушивания stdin канала. Все данные полученные от пользователя собираются в структуру данных полученноя из генерации прото файла и сохранения в файл указанный в аргументе при запуске. Написан небольшой тест.

    _По сути реализовано_:
  * Чтение файли и унмаршалинг в структуру
  * Интерактивный опрос в консоли с заполнением записной книжки
  * Маршелинг или преобразование из структуры данных в байты и запись в файл.

* Во второй части (list_people.go) реализовано:
  * Чтение из файла
  * Преобразование байтов в структуру (унмаршелинг)
  * Отображение в консоли.

# Little about protoc-gen-go

[About protoc-gen-go](https://protobuf.dev/reference/go/go-generated/) its options and using general.

Where in the output directory the generated .pb.go file is placed depends on the compiler flags. There are several output modes:

* `paths=import`
* `module=$PREFIX`
* `paths=source_relative`

# References

- [source](https://protobuf.dev/getting-started/gotutorial/) where explanation all steps.
