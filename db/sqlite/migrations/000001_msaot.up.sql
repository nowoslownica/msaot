create table grammar_positions (
  id int primary key,
  gCase int, -- 1: nom, 2: gen, 3: acc, 4: dat, 5: instr, 6: loc, 7: voc, 8: part, 9: prep
  gPerson int, -- 1-2-3
  gNumber int, -- 1-2-3
  gTense int, -- 1-2: present, 3-6: past, 7-8: future
  gGender int -- 1: f, 2: m, 3: n
);

create table lemmas (
  id int primary key,
  value text not null,
  pos text not null,
  changeSchema int -- 1-10: noun, 11-20: verb, 21-30: pronoun
);

create table flexies (
  id int primary key,
  value text not null,
  lemmaId int references lemmas (id) on delete restrict,
  gPositionId int references grammar_positions (id) on delete restrict
);

insert into grammar_positions values(0, null, null, null, null, null);

-- declension noun/adjective
insert into grammar_positions values(1, 1, null, 1, null, 1);
insert into grammar_positions values(2, 2, null, 1, null, 1);
insert into grammar_positions values(3, 3, null, 1, null, 1);
insert into grammar_positions values(4, 4, null, 1, null, 1);
insert into grammar_positions values(5, 5, null, 1, null, 1);
insert into grammar_positions values(6, 6, null, 1, null, 1);
insert into grammar_positions values(7, 7, null, 1, null, 1);
insert into grammar_positions values(8, 1, null, 3, null, 1);
insert into grammar_positions values(9, 2, null, 3, null, 1);
insert into grammar_positions values(10, 3, null, 3, null, 1);
insert into grammar_positions values(11, 4, null, 3, null, 1);
insert into grammar_positions values(12, 5, null, 3, null, 1);
insert into grammar_positions values(13, 6, null, 3, null, 1);
insert into grammar_positions values(14, 7, null, 3, null, 1);

-- conjugation verb
-- present
insert into grammar_positions values(15, null, 1, 1, 1, null);
insert into grammar_positions values(16, null, 2, 1, 1, null);
insert into grammar_positions values(17, null, 3, 1, 1, null);
insert into grammar_positions values(18, null, 1, 3, 1, null);
insert into grammar_positions values(19, null, 2, 3, 1, null);
insert into grammar_positions values(20, null, 3, 3, 1, null);
-- aorist
insert into grammar_positions values(21, null, 1, 1, 3, null);
insert into grammar_positions values(22, null, 2, 1, 3, null);
insert into grammar_positions values(23, null, 3, 1, 3, null);
insert into grammar_positions values(24, null, 1, 3, 3, null);
insert into grammar_positions values(25, null, 2, 3, 3, null);
insert into grammar_positions values(26, null, 3, 3, 3, null);


-- participle
insert into grammar_positions values(27, null, null, 1, 3, 1);
insert into grammar_positions values(28, null, null, 1, 3, 2);
insert into grammar_positions values(29, null, null, 1, 3, 3);
