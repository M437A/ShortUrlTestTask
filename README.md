# ShortUrlTestTask

Программа создает короткий url адрес, через который можно получить доступ к длинной ссылке,
короктий url адрес можно ввести либо в поисковую строку, либо в специальное поле,
данные можно сохранить либо в базу данных, либо в память программы, 
для выбора можно использовать флаг "-d", с его помощью ссылки сохраняться в базе данных

Перед запуском программы не забудьте подключиться к совей базе данных
Для этого в файле createSQL.go введите данные вашей таблицы, а также убедитесь, что в функциях ниже 
имя таблицы и название ее ячеек совпадает с вашим
