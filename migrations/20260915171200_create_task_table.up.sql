-- 1. Включаем расширение для генерации UUID (если еще не включено)
-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- 2. Создаем общую функцию для триггера updated_at (если ее еще нет в БД)
CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 3. Создаем таблицу tasks с CHECK-ограничениями на основе констант Go
CREATE TABLE tasks (
    id          UUID PRIMARY KEY DEFAULT uuidv7(),
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    status      VARCHAR(50) NOT NULL,
    type        VARCHAR(50) NOT NULL,
    description TEXT NOT NULL DEFAULT '',

    -- Строгая валидация TaskStatus
    CONSTRAINT chk_task_status CHECK (status IN (
        'new', 'prepaid', 'assigned', 'process', 'done', 'confirm', 'finish', 'cancel'
    )),
    
    -- Строгая валидация TaskType
    CONSTRAINT chk_task_type CHECK (type IN (
        'mark', 'verify', 'classify', 'compare', 'take_photo', 'take_video', 'take_audio', 'location', 'deliver', 'other'
    ))
);

-- 4. Создаем триггер для автоматического обновления поля updated_at
CREATE TRIGGER update_tasks_modtime
    BEFORE UPDATE ON tasks
    FOR EACH ROW
    EXECUTE FUNCTION update_modified_column();
