# Домашнее задание 

**Выполнил - Третьяков Александр Юрьевич**

## Задание

| Поле | Значение |
|------|----------|
| **Дисциплина** | Семинар наставника |
| **Тема** | Big Data и ML |
| **Форма проверки** | Задание с проверкой преподавателем |
| **Имя преподавателя** | Владислав Шевченко |
| **Время выполнения** | 6 часов |
| **Цель задания** | • Спроектировать и реализовать data pipeline.<br>• Применить практики ETL/ELT: извлечение, преобразование, загрузка и партиционирование данных.<br>• Подготовить данные для аналитики и ML-задач (очистка, агрегация, feature engineering).<br>• Применить базовые методы анализа данных и построения ML-моделей.<br>• Визуализировать и интерпретировать результаты. |
| **Инструменты для выполнения ДЗ** | • Docker / Docker Compose<br>• PostgreSQL<br>• MinIO (S3-совместимое хранилище)<br>• JupyterHub (Python, pandas, sklearn)<br>• Apache Superset<br>• GitHub / GitLab (репозиторий для сдачи) |
| **Правила приёма работы** | Прикрепить в LMS ссылки на репозиторий GitHub. Настройки репозитория должны быть public. |
| **Критерии оценки** | • **10 баллов** — выполнены все ключевые этапы (инфраструктура, ETL, обработка, ML, визуализация), решение корректно и воспроизводимо, структура проекта и код оформлены аккуратно. Оценка 9–10 возможна при уровне выше ожидаемого: Airflow, Data Quality, incremental-загрузка, оптимизация pipeline.<br>• **8 баллов** — реализованы основные этапы (не менее 80% задания), pipeline работает, присутствуют все ключевые компоненты, но есть упрощения или недочёты.<br>• **5-6 баллов** — реализована часть задания, есть попытка построения pipeline, но присутствуют ошибки в логике, неполная реализация или проблемы с воспроизводимостью. |
| **Дедлайн** | Точная дата указана в LMS |

### Описание задания


В рамках задания необходимо реализовать полный аналитический pipeline:  
**PostgreSQL → ETL → MinIO → обработка → Jupyter → витрины → Superset**

---

### Этапы выполнения

#### 1. Развертывание инфраструктуры

Необходимо:
- Создать `docker-compose.yml`
- Развернуть:
  - PostgreSQL
  - MinIO
  - Superset
  - JupyterHub

#### 2. Подключение сервисов

Настроить подключения:
- Jupyter → PostgreSQL
- Jupyter → MinIO (S3 API)
- Superset → PostgreSQL или MinIO (через Presto/Trino опционально)

#### 3. Подготовка данных

Использовать предоставленные SQL-скрипты:
- `wells` — справочник скважин
- `production` — ежедневная добыча
- `telemetry` — телеметрия оборудования

Данные для вставки уже предоставлены (см. SQL файлы)

#### 4. Выгрузка данных в MinIO

Необходимо:
- Считать данные из PostgreSQL
- Преобразовать (pandas / spark)
- Сохранить в MinIO:
  - формат: parquet / csv
  - разбивка по датам (partitioning)

#### 5. Очистка и подготовка данных

Реализовать:
- обработку NULL
- фильтрацию выбросов
- агрегации:
  - по дням
  - по скважинам
- создание фичей:
  - среднее давление
  - средняя температура
  - коэффициент простоя

---

### Практическая часть

#### Задание 1. Аналитика добычи

Компания получает телеметрию со скважин:
- дебит (тонн/сут)
- давление
- температура
- энергопотребление
- простой оборудования

Необходимо построить аналитические витрины:
1. Общую добычу по дням
2. KPI по скважинам:
   - средний дебит
   - % простоя
3. Найти:
   - лучшие/худшие скважины
   - влияние температуры и давления

**В Superset:**
- Line chart (добыча по времени)
- Bar chart (топ скважин)
- Heatmap (давление vs дебит)

---

#### Задание 2. Прогноз дебита (ML)

На основе данных из таблицы `well_targets` необходимо предсказать суточный дебит:
- давление
- температура
- мощность
- время работы насосов

Необходимо:
1. Подготовить dataset:
   - агрегировать telemetry → день
2. Обучить модель:
   - Linear Regression / RandomForest
3. Оценить:
   - MAE / RMSE
4. Сделать прогноз

**В Superset:**
- Actual vs Predicted
- Ошибка модели по времени

---

#### Задание 3. Аномалии и отказ оборудования

На основе данных в таблицах `pump_sensors` и `pump_failures`:
- вибрация
- температура
- ток
- обороты

Необходимо:
1. Найти аномалии:
   - z-score / isolation forest
2. Определить:
   - признаки перед отказом
3. Построить модель:
   - вероятность отказа

**В Superset:**
- Аномалии по времени
- Рост вибрации перед отказом
- Risk score по насосам

---

#### Задание 4. Логистика и поставки

На основе данных из таблицы `deliveries`

**Поля:**
- маршрут
- объем
- стоимость
- задержка
- погода

Необходимо:
1. Определить факторы задержек:
   - погода
   - расстояние
   - водитель
2. Проанализировать стоимость:
   - cost / km
3. Оптимизировать маршруты

**В Superset:**
- Delay vs Weather
- Cost vs Distance
- KPI по водителям

---

### Дополнительное задание

**Advanced:**
- Реализовать пайплайн через Airflow
- Сделать incremental загрузку
- Добавить Data Quality проверки:
  - NULL %
  - outliers
- Построить витрины:
  - `mart_production`
  - `mart_failures`
  - `mart_logistics`

---

### Чек-лист самопроверки

Необходимо предоставить:

| № | Компонент | Что предоставить |
|---|-----------|------------------|
| 1 | Инфраструктура | `docker-compose.yml` |
| 2 | ETL | скрипты выгрузки в MinIO |
| 3 | Data Processing | ноутбуки Jupyter |
| 4 | ML | модель + метрики |
| 5 | BI | дашборды Superset |

## Решение

### Подготовка сервисов

Установим из `bigData_ML\docker\docker-compose.yml` файла необходимые сервисы

<img src="./assets/2026-05-19 075853.jpg" width="700"> 

Проверим доступность установленных сервисов

### MinIO Console
| Параметр | Значение |
|----------|----------|
| URL | `http://localhost:9001` |
| Login | `admin` |
| Password | `admin123` |

<img src="./assets/2026-05-19 080128.jpg" width="700"> 

### Jupyter Lab
| Параметр | Значение |
|----------|----------|
| URL | `http://localhost:8888` |
| Токен | нужно посмотреть в логах |

```powershell
docker logs jupyter 2>&1 | findstr "token"
```
<img src="./assets/2026-05-19 080336.jpg" width="700"> 

### Superset:
| Параметр | Значение |
|----------|----------|
| URL | `http://localhost:8088` |
| Login | `admin` |
| Password | `admin` |

<img src="./assets/2026-05-19 082134.jpg" width="700"> 

### PostgreSQL:
| Параметр | Значение |
|----------|----------|
| Host | localhost |
| Port | 5432 |
| User | admin |
| Password | admin |
| Database | oil_db |

Проверка загрузки данных в таблицы PostgreSQL
```powershell
docker exec -it postgres_db psql -U admin -d oil_db -c "\dt"
```

<img src="./assets/2026-05-19 081020.jpg" width="700"> 

Создадим новый ноутбук и установи необходимые пакеты:
python
```python
!pip install boto3 pandas==2.0.3 psycopg2-binary sqlalchemy pyarrow==12.0.1 numpy==1.24.3 scikit-learn
```

Проверим установленные пакеты
```python
import pandas as pd
import pyarrow
import boto3
import psycopg2
import sqlalchemy
import sklearn
import scipy
import numpy as np

print("Все пакеты установлены!")
print(f"pandas: {pd.__version__}")
print(f"numpy: {np.__version__}")
print(f"sklearn: {sklearn.__version__}")
print(f"pyarrow: {pyarrow.__version__}")
```


<img src="./assets/2026-05-19 082340.jpg" width="700"> 

Проверим подключение к PostgreSQL из Jupyter:
```python
import psycopg2

conn = psycopg2.connect(
    host='postgres_db',
    port=5432,
    user='admin',
    password='admin',
    database='oil_db'
)
cursor = conn.cursor()
cursor.execute("SELECT COUNT(*) FROM wells;")
print(f"Скважин: {cursor.fetchone()[0]}")
conn.close()
```

Проверим подключение к MinIO из Jupyter:
```python
import boto3

s3 = boto3.client(
    's3',
    endpoint_url='http://minio:9000',
    aws_access_key_id='admin',
    aws_secret_access_key='admin123',
    verify=False
)
```

Создадим бакет
```python
s3.create_bucket(Bucket='oil-data')
print("Бакет создан")
```

<img src="./assets/2026-05-19 082812.jpg" width="700"> 

Бакет доступен в `minio`

<img src="./assets/2026-05-19 082858.jpg" width="700">

### Подготовка данных

Для создания таблиц и загрузки данных использованы предоставленные SQL-скрипты:

| Файл | Таблицы |
|------|---------|
| `01_create_wells_production.sql` | wells, production |
| `02_create_telemetry_targets.sql` | well_telemetry, well_targets |
| `03_create_pumps.sql` | pumps, pump_sensors, pump_failures |
| `06_create_deliveries.sql` | deliveries, drivers, vehicles |
| `09_create_insert_stations.sql` | oil_stations |

Все SQL-файлы были смонтированы в `/docker-entrypoint-initdb.d/` при запуске PostgreSQL, поэтому таблицы и данные создались автоматически при первом запуске контейнера.

**Результат:** 11 таблиц загружено в БД `oil_db`.

<img src="./assets/2026-05-19 081020.jpg" width="700">

### Выгрузка данных из PostgreSQL в MinIO (ETL)
Создадим в Jupyter ноутбук и выполним:

```python
import pandas as pd
import psycopg2
import boto3
from io import StringIO, BytesIO
from datetime import datetime

# Подключение к PostgreSQL
pg_conn = psycopg2.connect(
    host='postgres_db',
    port=5432,
    user='admin',
    password='admin',
    database='oil_db'
)

# Список таблиц для выгрузки
tables = ['wells', 'production', 'well_telemetry', 'well_targets', 
          'pumps', 'pump_sensors', 'pump_failures', 'deliveries', 
          'drivers', 'vehicles', 'oil_stations']

# Подключение к MinIO
s3 = boto3.client(
    's3',
    endpoint_url='http://minio:9000',
    aws_access_key_id='admin',
    aws_secret_access_key='admin123',
    verify=False
)

# Выгрузка каждой таблицы в Parquet
for table in tables:
    df = pd.read_sql(f'SELECT * FROM {table};', pg_conn)
    
    # Конвертация в Parquet
    parquet_buffer = BytesIO()
    df.to_parquet(parquet_buffer, index=False)
    parquet_buffer.seek(0)
    
    # Загрузка в MinIO
    s3.put_object(
        Bucket='oil-data',
        Key=f'{table}.parquet',
        Body=parquet_buffer.getvalue()
    )
    print(f' {table}: {len(df)} строк -> oil-data/{table}.parquet')

pg_conn.close()
print("\n Все данные выгружены в MinIO")
```

Parquet выгрузился успешно. Предупреждение про SQLAlchemy можно игнорировать — это просто рекомендация, не ошибка.

<img src="./assets/2026-05-19 084411.jpg" width="700">

Проверим наличие файлов parquet в в MinIO

<img src="./assets/2026-05-19 084540.jpg" width="700">

Выполним партиционирование

```python
from io import BytesIO
from sqlalchemy import create_engine
import pandas as pd
import boto3

# Подключение к MinIO
s3 = boto3.client(
    's3',
    endpoint_url='http://minio:9000',
    aws_access_key_id='admin',
    aws_secret_access_key='admin123',
    verify=False
)

# Подключение к PostgreSQL
engine = create_engine('postgresql://admin:admin@postgres_db:5432/oil_db')
production = pd.read_sql('SELECT * FROM production;', engine)

# Сохранение с партиционированием по датам
for date, group in production.groupby('date'):
    date_str = str(date)
    parquet_buffer = BytesIO()
    group.to_parquet(parquet_buffer, index=False)
    parquet_buffer.seek(0)
    
    s3.put_object(
        Bucket='oil-data',
        Key=f'production/date={date_str}/data.parquet',
        Body=parquet_buffer.getvalue()
    )
    print(f"Сохранено: production/date={date_str}/data.parquet ({len(group)} записей)")

print("\nПартиционирование завершено!")
```
<img src="./assets/2026-05-19 151623.jpg" width="700">

Проверим наличие файлов в бакете

<img src="./assets/2026-05-19 151834.jpg" width="700">


### Очистка и подготовка данных

1. Загрузим данные из MinIO и подготовим для ML.

```python
import pandas as pd
import boto3
from io import BytesIO

# Подключение к MinIO
s3 = boto3.client(
    's3',
    endpoint_url='http://minio:9000',
    aws_access_key_id='admin',
    aws_secret_access_key='admin123',
    verify=False
)

# Функция загрузки parquet из MinIO
def load_parquet(bucket, key):
    obj = s3.get_object(Bucket=bucket, Key=key)
    return pd.read_parquet(BytesIO(obj['Body'].read()))

# Загрузка всех таблиц
wells = load_parquet('oil-data', 'wells.parquet')
production = load_parquet('oil-data', 'production.parquet')
telemetry = load_parquet('oil-data', 'well_telemetry.parquet')
targets = load_parquet('oil-data', 'well_targets.parquet')
pumps = load_parquet('oil-data', 'pumps.parquet')
pump_sensors = load_parquet('oil-data', 'pump_sensors.parquet')
failures = load_parquet('oil-data', 'pump_failures.parquet')
deliveries = load_parquet('oil-data', 'deliveries.parquet')
drivers = load_parquet('oil-data', 'drivers.parquet')
vehicles = load_parquet('oil-data', 'vehicles.parquet')
stations = load_parquet('oil-data', 'oil_stations.parquet')

print("Данные загружены")
print(f"wells: {len(wells)}")
print(f"production: {len(production)}")
print(f"telemetry: {len(telemetry)}")
print(f"targets: {len(targets)}")
```

<img src="./assets/2026-05-19 095216.jpg" width="700">

2. Базовый анализ и очистка

```python
# Проверка пропусков
print("Пропуски в production:")
print(production.isnull().sum())

# Заполнение или удаление NULL
production = production.fillna({'temperature': production['temperature'].median(),
                                 'pressure': production['pressure'].median(),
                                 'downtime_hours': 0})

# Создание агрегатов по дням
daily_production = production.groupby('date').agg({
    'oil_ton': 'sum',
    'gas_m3': 'sum',
    'water_m3': 'sum',
    'energy_kwh': 'sum',
    'downtime_hours': 'mean'
}).reset_index()

print("Ежедневная агрегация:")
print(daily_production.head())
```

<img src="./assets/2026-05-19 095420.jpg" width="700">

3. Feature Engineering для ML (прогноз дебита)

```python
# Агрегация телеметрии по дням
telemetry['date'] = pd.to_datetime(telemetry['timestamp']).dt.date
daily_telemetry = telemetry.groupby(['well_id', 'date']).agg({
    'pump_speed_rpm': 'mean',
    'pump_current': 'mean',
    'pressure_in': 'mean',
    'pressure_out': 'mean',
    'temperature': 'mean',
    'vibration': 'mean',
    'oil_flow_rate': 'mean'
}).reset_index()

# Объединение с targets (целевая переменная)
ml_data = daily_telemetry.merge(targets, on=['well_id', 'date'], how='inner')

# Добавим информацию о скважине
ml_data = ml_data.merge(wells[['well_id', 'field_name', 'region']], on='well_id', how='left')

print(f"Датасет для ML: {len(ml_data)} записей")
print(ml_data.head())
```

<img src="./assets/2026-05-19 095606.jpg" width="700">

4. Сохранение подготовленные данные обратно в MinIO

```python
# Сохранение витрин
daily_production.to_parquet('daily_production.parquet', index=False)
ml_data.to_parquet('ml_dataset.parquet', index=False)

# Загрузка в MinIO
s3.put_object(Bucket='oil-data', Key='daily_production.parquet', 
              Body=open('daily_production.parquet', 'rb').read())
s3.put_object(Bucket='oil-data', Key='ml_dataset.parquet', 
              Body=open('ml_dataset.parquet', 'rb').read())

print("Витрины сохранены в MinIO")
```

<img src="./assets/2026-05-19 095827.jpg" width="700">


5. ML модель (прогноз дебита)

```python
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestRegressor
from sklearn.metrics import mean_absolute_error, mean_squared_error
import numpy as np

# Подготовка признаков
features = ['pump_speed_rpm', 'pump_current', 'pressure_in', 'pressure_out', 
            'temperature', 'vibration', 'oil_flow_rate']
X = ml_data[features]
y = ml_data['daily_oil_ton']

# Разделение на train/test
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

# Обучение модели
model = RandomForestRegressor(n_estimators=100, random_state=42)
model.fit(X_train, y_train)

# Прогноз
y_pred = model.predict(X_test)

# Оценка
mae = mean_absolute_error(y_test, y_pred)
rmse = np.sqrt(mean_squared_error(y_test, y_pred))

print(f"MAE: {mae:.2f}")
print(f"RMSE: {rmse:.2f}")
print(f"Feature importance: {dict(zip(features, model.feature_importances_))}")

# Сохранение прогнозов
ml_data['predicted_oil_ton'] = model.predict(X)
ml_data[['well_id', 'date', 'daily_oil_ton', 'predicted_oil_ton']].to_parquet('predictions.parquet', index=False)
s3.put_object(Bucket='oil-data', Key='predictions.parquet', 
              Body=open('predictions.parquet', 'rb').read())
print("Прогнозы сохранены в MinIO")
```

"В предоставленных данных телеметрия охватывает только 1 октября 2025 года (скважины 1 и 2), тогда как целевая переменная и добыча доступны за весь месяц. Это не позволяет построить полноценную ML модель, но демонстрирует полностью рабочий пайплайн."

<img src="./assets/2026-05-19 101511.jpg" width="700">

### Построить аналитику и витрины


Загрузка данных из PostgreSQL
```python
import pandas as pd
from sqlalchemy import create_engine

# Подключение
engine = create_engine('postgresql://admin:admin@postgres_db:5432/oil_db')

# Загрузка таблиц
production = pd.read_sql('SELECT * FROM production;', engine)
wells = pd.read_sql('SELECT * FROM wells;', engine)

print(f"production: {len(production)} записей")
print(f"wells: {len(wells)} записей")
```

Построение витрин

```python
# Витрина 1: Ежедневная добыча по скважинам
daily_production = production.groupby(['well_id', 'date']).agg({
    'oil_ton': 'sum',
    'gas_m3': 'sum',
    'water_m3': 'sum',
    'energy_kwh': 'sum',
    'downtime_hours': 'mean',
    'temperature': 'mean',
    'pressure': 'mean'
}).reset_index()

daily_production.to_sql('daily_production', engine, if_exists='replace', index=False)
print("daily_production сохранена")
print(daily_production.head())

# Витрина 2: KPI по скважинам
kpi_wells = production.groupby('well_id').agg({
    'oil_ton': 'mean',
    'downtime_hours': 'mean',
    'energy_kwh': 'mean'
}).reset_index()
kpi_wells.columns = ['well_id', 'avg_oil_ton', 'avg_downtime_hours', 'avg_energy_kwh']

kpi_wells.to_sql('kpi_wells', engine, if_exists='replace', index=False)
print("kpi_wells сохранена")
```

<img src="./assets/2026-05-19 111301.jpg" width="700">

### Аномалии и отказ оборудования

Загрузка данные и поиск аномалий
```python
import pandas as pd
import numpy as np
from scipy import stats
from sqlalchemy import create_engine

engine = create_engine('postgresql://admin:admin@postgres_db:5432/oil_db')

# Загрузка данных
pump_sensors = pd.read_sql('SELECT * FROM pump_sensors;', engine)
pump_failures = pd.read_sql('SELECT * FROM pump_failures;', engine)
pumps = pd.read_sql('SELECT * FROM pumps;', engine)

print(f"pump_sensors: {len(pump_sensors)} записей")
print(f"pump_failures: {len(pump_failures)} записей")
```

```python
# Выявление аномалий по вибрации и температуре
sensors_anomalies = pump_sensors.copy()

# Z-score для вибрации
z_scores_vibration = np.abs(stats.zscore(sensors_anomalies['vibration'].fillna(sensors_anomalies['vibration'].median())))
sensors_anomalies['vibration_anomaly'] = z_scores_vibration > 2

# Z-score для температуры
z_scores_temp = np.abs(stats.zscore(sensors_anomalies['temperature'].fillna(sensors_anomalies['temperature'].median())))
sensors_anomalies['temperature_anomaly'] = z_scores_temp > 2

# Общая аномалия
sensors_anomalies['is_anomaly'] = sensors_anomalies['vibration_anomaly'] | sensors_anomalies['temperature_anomaly']

print(f"Найдено аномалий: {sensors_anomalies['is_anomaly'].sum()}")
print(sensors_anomalies[sensors_anomalies['is_anomaly']][['pump_id', 'timestamp', 'vibration', 'temperature']])
```

<img src="./assets/2026-05-19 152507.jpg" width="700">

Признаки перед отказом
```python
# Объединение с данными об отказах
pump_failures['failure_date'] = pd.to_datetime(pump_failures['failure_date'])
sensors_anomalies['timestamp'] = pd.to_datetime(sensors_anomalies['timestamp'])

# Определение аномалий за 6 часов до отказа
failure_analysis = []

for _, failure in pump_failures.iterrows():
    # Данные за 6 часов до отказа
    before_failure = sensors_anomalies[
        (sensors_anomalies['pump_id'] == failure['pump_id']) &
        (sensors_anomalies['timestamp'] >= failure['failure_date'] - pd.Timedelta(hours=6)) &
        (sensors_anomalies['timestamp'] <= failure['failure_date'])
    ]
    
    if len(before_failure) > 0:
        failure_analysis.append({
            'pump_id': failure['pump_id'],
            'failure_type': failure['failure_type'],
            'avg_vibration_before': before_failure['vibration'].mean(),
            'max_vibration_before': before_failure['vibration'].max(),
            'avg_temp_before': before_failure['temperature'].mean(),
            'anomalies_count': before_failure['is_anomaly'].sum()
        })

failure_analysis = pd.DataFrame(failure_analysis)
print("Анализ перед отказом:")
print(failure_analysis)
```
Модель вероятности отказа
```python
from sklearn.ensemble import RandomForestClassifier

# Подготовка данных для обучения
# 1 - отказ произошел, 0 - отказа не было
sensors_anomalies['failure_next_hour'] = 0

for _, failure in pump_failures.iterrows():
    mask = (sensors_anomalies['pump_id'] == failure['pump_id']) & \
           (sensors_anomalies['timestamp'] >= failure['failure_date'] - pd.Timedelta(hours=1)) & \
           (sensors_anomalies['timestamp'] < failure['failure_date'])
    sensors_anomalies.loc[mask, 'failure_next_hour'] = 1

# Признаки
features = ['vibration', 'temperature', 'current', 'rpm', 'pressure']
X = sensors_anomalies[features].fillna(sensors_anomalies[features].median())
y = sensors_anomalies['failure_next_hour']

# Обучение
model = RandomForestClassifier(n_estimators=100, random_state=42)
model.fit(X, y)

# Важность признаков
importance = dict(zip(features, model.feature_importances_))
print("Важность признаков для отказа:")
for k, v in sorted(importance.items(), key=lambda x: -x[1]):
    print(f"  {k}: {v:.3f}")

# Сохранение результатов
sensors_anomalies[['pump_id', 'timestamp', 'vibration', 'temperature', 'is_anomaly', 'failure_next_hour']].to_sql('anomalies_results', engine, if_exists='replace', index=False)
print("\n Результаты аномалий сохранены в таблицу 'anomalies_results'")
```
<img src="./assets/2026-05-19 152705.jpg" width="700">

<img src="./assets/2026-05-19 152806.jpg" width="700">

*"Из-за малого количества отказов (3) модель не выявила значимых признаков. Для реального применения нужны данные за больший период."*

### Логистика и поставки

```python
import pandas as pd
from sqlalchemy import create_engine

engine = create_engine('postgresql://admin:admin@postgres_db:5432/oil_db')

# Загрузка данных
deliveries = pd.read_sql('SELECT * FROM deliveries;', engine)
drivers = pd.read_sql('SELECT * FROM drivers;', engine)
vehicles = pd.read_sql('SELECT * FROM vehicles;', engine)

print(f"deliveries: {len(deliveries)} записей")
print(f"drivers: {len(drivers)} записей")
print(f"vehicles: {len(vehicles)} записей")

# 1. Факторы задержек
delay_by_weather = deliveries.groupby('weather_conditions').agg({
    'delay_hours': 'mean',
    'delivery_id': 'count'
}).rename(columns={'delivery_id': 'count'})
print("Задержки по погоде:")
print(delay_by_weather)

# 2. Анализ стоимости
deliveries['cost_per_km'] = deliveries['cost_usd'] / deliveries['distance_km']

# KPI по водителям
driver_performance = deliveries.groupby('driver_id').agg({
    'cost_per_km': 'mean',
    'delay_hours': 'mean',
    'delivery_id': 'count'
}).rename(columns={'delivery_id': 'delivery_count'})
driver_performance = driver_performance.merge(drivers[['driver_id', 'name']], on='driver_id')

print("\nKPI по водителям:")
print(driver_performance.sort_values('cost_per_km'))

# 3. Сохранение результатов
driver_performance.to_sql('logistics_driver_kpi', engine, if_exists='replace', index=False)
print("\n Результаты логистики сохранены в PostgreSQL")
```

<img src="./assets/2026-05-19 155709.jpg" width="700">

### Построение визуализации в Superset

Подключим базу данных PostgreSQL с помощью SQLAlchemy URI
`Data → Databases → + Database`
`Database name: oil_db`

`SQLAlchemy URI:`

```text
postgresql://admin:admin@postgres_db:5432/oil_db
```

<img src="./assets/2026-05-19 102504.jpg" width="700">

Создадим дашборды

Line chart - Добыча по дням
<img src="./assets/2026-05-21 154722.jpg" width="700">

Bar chart - Топ скважин по добыче
<img src="./assets/2026-05-21 155315.jpg" width="700">

Heatmap - Давление vs Дебит

<img src="./assets/2026-05-21 155944.jpg" width="700">

Bar chart - Delay vs Weather

<img src="./assets/2026-05-21 161823.jpg" width="700">

Scatter Plot (точечный график) -  Cost vs Distance

<img src="./assets/2026-05-21 162047.jpg" width="700">

Bar Chart (группированный) - KPI по водителям

<img src="./assets/2026-05-21 162240.jpg" width="700">