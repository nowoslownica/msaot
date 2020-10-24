create table grammar_positions (
  id int primary key,
  gCase int,
  gPerson int,
  gNumber int,
  gTense int,
  gGender int,
  declension int,
  conjugation int,
)

create table flexies (
  id int primary key,
  value string not null,
  normal string not null,
  pos string not null,
  gPosition int references grammar_positions (id) on delete restrict
);

-- declension 1A
insert into grammar_positions values(1, 1, null, 1, null, 1, 1, null);
insert into grammar_positions values(2, 2, null, 1, null, 1, 1, null);
insert into grammar_positions values(3, 3, null, 1, null, 1, 1, null);
insert into grammar_positions values(4, 4, null, 1, null, 1, 1, null);
insert into grammar_positions values(5, 5, null, 1, null, 1, 1, null);
insert into grammar_positions values(6, 6, null, 1, null, 1, 1, null);
insert into grammar_positions values(7, 7, null, 1, null, 1, 1, null);
insert into grammar_positions values(8, 1, null, 2, null, 1, 1, null);
insert into grammar_positions values(9, 2, null, 2, null, 1, 1, null);
insert into grammar_positions values(10, 3, null, 2, null, 1, 1, null);
insert into grammar_positions values(11, 4, null, 2, null, 1, 1, null);
insert into grammar_positions values(12, 5, null, 2, null, 1, 1, null);
insert into grammar_positions values(13, 6, null, 2, null, 1, 1, null);
insert into grammar_positions values(14, 7, null, 2, null, 1, 1, null);

insert into flexies values(1, 'мама', 'мама', 'n', 1)
insert into flexies values(1, 'мамы', 'мама', 'n', 2)
insert into flexies values(1, 'маму', 'мама', 'n', 3)
insert into flexies values(1, 'маме', 'мама', 'n', 4)
insert into flexies values(1, 'мамой', 'мама', 'n', 5)
insert into flexies values(1, 'маме', 'мама', 'n', 6)
insert into flexies values(1, 'маме', 'мама', 'n', 7)
insert into flexies values(1, 'мамо', 'мама', 'n', 8)
insert into flexies values(1, 'мамы', 'мама', 'n', 9)
insert into flexies values(1, 'мам', 'мама', 'n', 10)
insert into flexies values(1, 'мамам', 'мама', 'n', 11)
insert into flexies values(1, 'мамам', 'мама', 'n', 12)
insert into flexies values(1, 'мамах', 'мама', 'n', 13)
insert into flexies values(1, 'мамах', 'мама', 'n', 14)
insert into flexies values(1, 'мамы', 'мама', 'n', 15)
-- declension 1B
-- insert into grammar_positions values(15, 1, null, 1, null, 1, 1, null);
-- insert into grammar_positions values(16, 2, null, 1, null, 1, 1, null);
-- insert into grammar_positions values(17, 3, null, 1, null, 1, 1, null);
-- insert into grammar_positions values(18, 4, null, 1, null, 1, 1, null);
-- insert into grammar_positions values(19, 5, null, 1, null, 1, 1, null);
-- insert into grammar_positions values(20, 6, null, 1, null, 1, 1, null);
-- insert into grammar_positions values(21, 7, null, 1, null, 1, 1, null);