INSERT INTO tbl_user (username, password, email, profile) VALUES 
---username		password	email				profile text
('cora23',		'qwerty',	'mikemcgrat11@gmail.com',	'Настоящие бедствия переносятся легче, нежели те, которые угрожают нам в будущем.'),

('giletobserver', 	'qwertyy', 	'giletobserver@mail.com',	'Русский поэт, эссеист, переводчик'),

('meeksat', 		'qwerty',	'meeksat@gmail.com',		'Когда хвалят глаза, то это значит, что остальное никуда не годится'),

('luge',		'qwerty',	'luge@gmail.com',		'российская певица, видеоблогер, автор песен, телеведущая.'),

('spatter',		'qwerty',	'spatter@gmail.com',		'известный советский астроном, астрофизик, член-корреспондент АН СССР');



INSERT INTO tbl_post (title, content, author_id) VALUES 

('Узбекистан мог потерять часть запасов лука из-за морозов и начать его импортировать', 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry`s', 1), 

('Шахматист Нодирбек Абдусатторов сенсационно победил Магнуса Карлсена','На супертурнире Tata Steel Chess гроссмейстер Нодирбек Абдусатторов одержал сенсационную победу над действующим чемпионом мира по классическим шахматам Магнусом Карлсеном.',1),

('Как победить энергокризис в Узбекистане. Мнение эксперта','Узбекистан, Ташкент – АН Podrobno.uz. Аномальные холода, пришедшие в страну этой зимой, в очередной раз подчеркнули проблемы, которые копились долгие годы в энергетической системе Узбекистана. Газа не хватает, в стране вводятся отключения электричества, газовые заправки работают считаные часы. Изменится ли к лучшему эта ситуация в будущем году однозначно сказать сложно. Скорее всего, при текущем отношении чиновников к проблеме все может только усугубиться.',1),

('Чудесное спасение. Как трехногая умирающая собака из Узбекистана получила шанс на лучшую жизнь в Италии','Узбекистан, Ташкент – АН Podrobno.uz. В Узбекистане в канун Нового года произошла настоящая сказка. Трехлапый умирающий пес получил шанс на новую жизнь, да ни где-нибудь, а в Италии. Эту трогательную историю спасения корреспонденту Podrobno.uz рассказала Ирода Маткаримова, основатель и руководитель приюта для бездомных собак "Хает".',2),

('Приведет ли сокращение министерств и госслужащих к уменьшению коррупции в Узбекистане. Мнение эксперта','Узбекистан, Ташкент – АН Podrobno.uz. Международный эксперт по борьбе с коррупцией Кодир Кулиев проанализировал стартовавшую административную реформу в Узбекистане на предмет ее влияния на уменьшение коррупции в госорганах. По его словам, большое количество министерств наносило ущерб общему госуправлению, потому что министерства судили о своей важности по тому, сколько они могут получить из госбюджета на траты. Чиновники соревновались внутри правительства за максимально возможную долю в общих расходах, независимо от общественного блага, которое они могли сделать с этими ресурсами.',2),

('Значительная часть опрошенных узбекистанцев не смогла ответить на вопрос о дружественности тех или иных стран – социологи СНГ','Узбекистан, Ташкент – АН Podrobno.uz. Социологи СНГ изучили общественное мнение своих стран в отношении России и ее программ по гуманитарному сотрудничеству. По результатам опроса 74% узбекистанцев положительно относятся к России. Она же является наиболее дружественной для 67% респондентов нашей страны. При этом интересно, что у нас в стране самая большая доля респондентов, которые не смогли ответить на вопрос о дружественности тех или иных стран, сообщает корреспондент Podrobno.uz.',2);



INSERT INTO tbl_tag (name) VALUES ('#энергетика'),('#Абдулла Арипов'),('#Очилбой Раматов'),('#Узбекистан'),('#Шавкат Мирзиёев'),('#жиноятчилик'),('#тиббиёт');



INSERT INTO post_tag VALUES (1,1),(1,2),(1,3), (2,1),(2,4),(2,5), (3,1),(3,6), (4,4),(4,5),(4,6), (5,6),(5,1),(5,7), (6,1),(6,5),(6,4);



INSERT INTO tbl_comment (content, create_time, author_id, post_id) VALUES
('this is a comment for second post', '2016-06-22 19:10:25-07', 1, 2),
('Статья просто классная, профессиональная. Как послушать, то таких классных идей исходит огромное количество, вот только с реализацией на практике большие проблемы. Как говорится хороших идей много, а телега стоит на одном месте.', '2023-01-22 19:10:25-07', 1, 1),
('"По рогам ему и промеж ему..." (В.С. ВЫСОЦКИЙ)','2023-01-22 19:10:25-07',1,3),

('Очередное малолетнее чмо, мнящее себя "вором в законе"','2023-01-20 19:10:25-07',2,1),
('Бои петухов, т.е. опущенных...и на этом зарабатывают.','2023-01-17 19:10:25-07',2,2),
('какие огромные деньги! Питух Еда, а Еда Сама Себя не зарубит вот они и пользуются Услугами Других Питухов дабы совершить таким образом Жертвоприношение,!','2023-01-16 19:10:25-07',2,4),

('А ещё есть статья - жестокое обращение с животными... нужно напомнить?','2023-01-16 19:10:25-07',3,5),
('Какое еще зрелище, наши предки этим веками занимались Нашли к чему придираться, как будто других проблем в нашей стране нету','2023-01-17 19:10:25-05',4,6),
('Предки много чем занимались.. например могли камнями забить женщину..в правовом (?) Государстве немного другие приоритеты..','2023-01-14 19:10:25-04',5,5);

---truncate tbl_post CASCADE2
