-- 1. Удаляем триггер
DROP TRIGGER IF EXISTS update_tasks_modtime ON tasks;

-- 2. Удаляем таблицу
DROP TABLE IF EXISTS tasks;
