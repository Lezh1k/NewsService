/*создаем единственную таблицу news.*/
create table news (
  id INT AUTO_INCREMENT,
  registered DATE NOT NULL,
  header TEXT,
  PRIMARY KEY(id)
);

/* скорей всего ещё нужно будет выбирать новости за определенную дату.
   вообще говоря обычно такие вещи выносятся в миграции, но из-за ограничений
   по времени пусть пока будет здесь
   TODO move to migrations*/
create index idx_news_registered on news(registered);
