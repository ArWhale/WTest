-- +goose Up
-- +goose NO TRANSACTION
-- DROP SCHEMA IF EXISTS public CASCADE;
-- CREATE SCHEMA public;
--
-- GRANT ALL ON SCHEMA public TO postgres;
-- GRANT ALL ON SCHEMA public TO public;

CREATE TABLE IF NOT EXISTS customers(
    id SERIAL PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    birthdate date NOT NULL,
    gender text NOT NULL,
    e_mail text NOT NULL,
    address text NOT NULL,
    UNIQUE(e_mail)
);

INSERT INTO customers(first_name, last_name, e_mail, gender, birthdate, address) VALUES
('Shanon', 'Lambert-Ciorwyn', 'slambertciorwyn3d@histats.com', 'Male', '2014-08-19', '079 Elgar Court'),
('Verena', 'Czajkowska', 'vczajkowska3e@comcast.net', 'Male', '2017-09-10', '920 Sommers Pass'),
('Aylmar', 'Pendle', 'apendle3f@fda.gov', 'Female', '2009-01-07', '22 Amoth Park'),
('Vanni', 'Firmin', 'vfirmin3g@wsj.com', 'Male', '2011-04-30', '46273 Dwight Park'),
('Hall', 'Mandy', 'hmandy3h@artisteer.com', 'Male', '2002-08-20', '48058 Donald Way'),
('Bronnie', 'Tenwick', 'btenwick3i@cnn.com', 'Male', '2009-11-07', '663 Donald Alley'),
('Hilarius', 'Verrills', 'hverrills3j@tuttocitta.it', 'Female', '1994-05-28', '104 Fallview Circle'),
('Janie', 'Quickfall', 'jquickfall3k@army.mil', 'Female', '2010-08-28', '5 Texas Avenue'),
('Cindi', 'O''Fihily', 'cofihily3l@foxnews.com', 'Male', '2003-10-10', '3660 Westerfield Parkway'),
('Barbi', 'Sisland', 'bsisland3m@apache.org', 'Female', '2015-03-16', '80 Vera Plaza'),
('Kristofer', 'Grahame', 'kgrahame3n@mapy.cz', 'Male', '2011-09-17', '55925 Miller Point'),
('Clotilda', 'Joseph', 'cjoseph3o@nih.gov', 'Male', '1998-07-11', '30706 Knutson Drive'),
('Peta', 'Wyon', 'pwyon3p@apache.org', 'Male', '1993-08-27', '69 1st Pass'),
('Cull', 'Hevner', 'chevner3q@google.fr', 'Female', '2015-10-03', '733 Thompson Court'),
('Zara', 'Daubeny', 'zdaubeny3r@washington.edu', 'Male', '2008-04-06', '5 Mallory Hill'),
('Becky', 'Blamey', 'bblamey3s@tumblr.com', 'Female', '2007-01-04', '748 Summerview Drive'),
('Nikita', 'Vreede', 'nvreede3t@liveinternet.ru', 'Male', '2005-07-01', '64 Forest Dale Point'),
('Julissa', 'Billing', 'jbilling3u@yale.edu', 'Male', '2005-04-10', '1 Fuller Point'),
('Christoforo', 'Tibald', 'ctibald3v@bigcartel.com', 'Male', '2009-02-25', '69068 Morrow Road'),
('Dacy', 'Ellse', 'dellse3w@slideshare.net', 'Male', '2016-04-05', '7697 Tony Alley'),
('Tanhya', 'Giamitti', 'tgiamitti3x@redcross.org', 'Male', '2020-01-08', '56689 Milwaukee Road'),
('Erich', 'Baskerfield', 'ebaskerfield3y@weibo.com', 'Male', '1995-10-20', '69 Grasskamp Park'),
('Marybeth', 'Crook', 'mcrook3z@patch.com', 'Female', '1993-10-21', '156 Mcbride Terrace'),
('Lizzy', 'Iannazzi', 'liannazzi40@cnet.com', 'Female', '2006-06-22', '5 Bellgrove Parkway'),
('Alina', 'Jakubovsky', 'ajakubovsky41@gov.uk', 'Male', '2007-05-26', '81 Hermina Plaza'),
('Sheffie', 'Ronchetti', 'sronchetti42@unblog.fr', 'Male', '1989-04-16', '46412 Eastwood Hill'),
('Golda', 'Summerfield', 'gsummerfield43@elpais.com', 'Female', '1994-10-12', '78095 Warrior Way'),
('Gae', 'Isac', 'gisac44@wp.com', 'Male', '1993-03-05', '16 Dixon Trail'),
('Frayda', 'de Clerk', 'fdeclerk45@seesaa.net', 'Female', '2021-02-15', '741 Sugar Drive') ON CONFLICT DO NOTHING;

-- +goose Down
