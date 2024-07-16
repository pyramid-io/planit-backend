-- Create user and grant privileges
CREATE USER 'planit'@'%' IDENTIFIED BY 'planit';
GRANT ALL PRIVILEGES ON planit.* TO 'planit'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;