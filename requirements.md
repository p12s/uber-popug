# Проектирование PopugJira

## 1. Собираем требования для PopugJira
![Общая схема](https://github.com/p12s/uber-popug/blob/master/domain-model-v1.png?raw=true)  
[Общая схема](https://lucid.app/lucidchart/408140b3-2eaf-4c6d-8006-dbde37212f06/edit?viewport_loc=164%2C9%2C1876%2C1136%2C0_0&invitationId=inv_02c3bab0-e84e-4117-a925-69861b830d41)    
  
###  Нефункциональные требования  
1. Никакого реалтайма тоже не нужно, хватит рефреша страницы для всех дашбордов и пользовательского интерфейса.  
2. Вопрос нагрузки не стоит, можно считать, что максимальное количество пользователей будет не больше 100 пользователей в минуту.  
  
### Функциональные требования  
В ходе обсуждения задачи с топ-менеджментом были выявлены следующие требования:    
    
#### Таск-трекер    
1. Таск-трекер должен быть отдельным дашбордом и доступен всем сотрудникам компании UberPopug Inc.    
**- Список тасков могут взять все попуги (query)**  
  
2. Авторизация в таск-трекере должна выполняться через общий сервис авторизации UberPopug Inc (у нас там инновационная система авторизации на основе формы клюва).    
**- Ввести свои идентиф. данные могут все попуги, в общем месте (query)**  
  
3. В таск-трекере должны быть только задачи. Проектов, скоупов и спринтов нет, потому что они не умещаются в голове попуга.    
**- Рядовые попуги могут просматривать только таски (query)**  
    
4. Новые таски может создавать кто угодно (администратор, начальник, разработчик, менеджер и любая другая роль). У задачи должны быть описание, статус (выполнена или нет) и попуг, на которого заассайнена задача.    
**- Все пупуги могут создавать таски (администратор, начальник, разработчик, менеджер и любая другая роль) (command)**    
**- Таск имеет описание, статус (выполнена или нет), и попуга, на которого заассайнена задача (дополнение к пред. требованию)**  
  
5. Менеджеры или администраторы должны иметь кнопку «заассайнить задачи», которая возьмёт все открытые задачи и рандомно заассайнит каждую на любого из сотрудников. Не успел закрыть задачу до реассайна — сорян, делай следующую.    
**- Менеджер и админ имеют возможность рандомно пере-ассайнить все открытые таски, каждому выпадает рандомное кол-во 0-10 (command)**    
    1. Ассайнить задачу можно на кого угодно, это может быть любой аккаунт из системы.  
    2. Ассайнить задачу можно только кнопкой «заассайнить задачи»  
    3. При нажатии кнопки «заассайнить задачи» все текущие не закрытые задачи должны быть случайным образом перетасованы между каждым аккаунтом в системе  
    4. Мы не заморачиваемся на ограничение по нажатию на кнопку «заассайнить задачи». Её можно нажимать хоть каждую секунду.  
    5. На одного сотрудника может выпасть любое количество новых задач, может выпасть ноль, а может и 10.  
  
6. Каждый сотрудник должен иметь возможность видеть в отдельном месте список заассайненных на него задач + отметить задачу выполненной.    
**- Все попуги могут видеть список своих задач (query)**    
**- Все попуги могут отметить свои задчи выполненными (command)**    
  
#### Аккаунтинг: кто сколько денег заработал  
1. Аккаунтинг должен быть в отдельном дашборде и доступным только для администраторов и бухгалтеров.  
    1. **Updated:** у обычных попугов доступ к аккаунтингу тоже должен быть. Но только к информации о собственных счетах (аудит лог + текущий баланс).   
    У админов и бухгалтеров должен быть доступ к общей статистике по деньгами заработанным (количество заработанных топ-менеджментом за сегодня денег + статистика по дням).    
**- В аккаунтинге админы и бухгалтеры могут видеть общую статистику по заработанным деньгами (количество заработанных топ-менеджментом за сегодня денег + статистика по дням) (query)**  
  
2. Авторизация в дешборде аккаунтинга должна выполняться через общий сервис аутентификации UberPopug Inc.    
**- shared Account, описано выше**  
  
3. У каждого из сотрудников должен быть свой счёт, который показывает, сколько за сегодня он получил денег. У счёта должен быть аудитлог того, за что были списаны или начислены деньги, с подробным описанием каждой из задач.    
**- Все попуги могут видеть информацию о собственном счете (аудитлог того, за что были списаны или начислены деньги, с подробным описанием каждой из задач + текущий баланс) (query)**  
  
4. Расценки:  
    - цены на задачу определяется единоразово, в момент появления в системе (можно с минимальной задержкой)  
    - **Updated:** цены рассчитывается без привязки к сотруднику  
    - **Updated:** формула, которая говорит сколько списать денег с сотрудника при ассайне задачи — `rand(-10..-20)$`  
    - **Updated:** формула, которая говорит сколько начислить денег сотруднику для выполненой задачи — `rand(20..40)$`  
    - деньги списываются сразу после ассайна на сотрудника, а начисляются после выполнения задачи.  
    - отрицательный баланс переносится на следующий день. Единственный способ его погасить - закрыть достаточное количество задач в течении дня.    
**- Цена на задачу определяется один раз, когда появляется в системе (можно с минимальной задержкой) (command)**    
**- При ассайне на попуга с него списывается стоимость таска (command)**    
**- После выполнения таска попуга начисляется рандоманя сумма (command)**    
  
5. Дешборд должен выводить количество заработанных топ-менеджментом за сегодня денег. **Updated:** удалилась строчка ~~Чтобы каждый попуг в компании знал, сколько заработал топ-менеджмент на сотрудниках~~.    
**- ? (query)**    
    1. т.е. сумма всех закрытых и созданных задач за день с противоположным знаком: `(sum(completed task amount) + sum(created task fee)) * -1`  
  
6. В конце дня необходимо:  
    a. считать сколько денег сотрудник получил за рабочий день  
    b. отправлять на почту сумму выплаты.    
**- Каждый сотрудник в конце дня должен получить инфо с суммой выплаты за прошедший день (command)**    
  
7. После выплаты баланса (в конце дня) он должен обнуляться, и в аудитлоге всех операций аккаунтинга должно быть отображено, что была выплачена сумма.    
**- Каждый сотрудник в конце дня получает обновление в аудитлог и обнуление баланса (command)**  

8. Дашборд должен выводить информацию по дням, а не за весь период сразу.    
**- ? (query)**    
    1. вообще хватит только за сегодня (всё равно попуги дальше не помнят), но если чувствуете, что успеете сделать аналитику за каждый день недели — будет круто  

#### Аналитика
1. Аналитика — это отдельный дашборд, доступный только админам.    
**- Аналитику могут посмотреть только админы-попуги (query)**    
2. Нужно указывать, сколько заработал топ-менеджмент за сегодня: сколько попугов ушло в минус.      
**- ? (query)**  
3. Нужно показывать самую дорогую задачу за день, неделю или месяц.  
    1. самой дорогой задачей является задача с наивысшей ценой из списка всех закрытых задач за определенный период времени  
    2. пример того, как это может выглядеть:  
        03.03 — самая дорогая задача — 28$    
        02.03 — самая дорогая задача — 38$  
        01.03 — самая дорогая задача — 23$  
        01-03 марта — самая дорогая задача — 38$      
**- ? (query)**  
  
## 2. Описываем query/comands  
- *Список тасков могут взять все попуги (query, read model)*  
- Ввести свои идентиф. данные могут все попуги, в общем месте (command) - (создание и возврат "ключа" - это изменение стейта)  
Actor   Account  
Command Login to app  
Data    Account (login, pass или форма клюва)  
Event   Account.Logined  
- *Рядовые попуги могут просматривать только таски (query, read model)*  
- Все пупуги могут создавать таски (администратор, начальник, разработчик, менеджер и любая другая роль, таск имеет описание, статус (выполнена или нет), и попуга, на которого заассайнена задача) (command)  
- Цена на задачу определяется один раз, когда появляется в системе (можно с минимальной задержкой) (тоже command, объединим с предыдущим действием. Если в будущем захотят вручную устанавливать цену - просто уберем рандом и добавим поле цены)  
Actor   Account  
Command Create task  
Data    Task (description, status, assigned account id + random price [rand(-10..-20)$]), Assigned account id  
Event   Task.Created  
- Менеджер и админ имеют возможность рандомно пере-ассайнить все открытые таски, каждому выпадает рандомное кол-во 0-10 (command)  
Actor   Account (with "manager" and "admin" roles)  
Command Random reassigne tasks  
Data    Tasks, Accounts (only without roles manager, admin)  
Event   Tasks.Reassigned  
- *Все попуги могут видеть список своих задач (query, read model)*  
- Все попуги могут отметить свои задчи выполненными (command)  
Actor   Account  
Command Complete task   
Data    Task  
Event   Task.Completed  
- *В аккаунтинге админы и бухгалтеры могут видеть общую статистику по заработанным деньгами (количество заработанных топ-менеджментом за сегодня денег + статистика по дням) (query, read model)*  
- *Все попуги могут видеть информацию о собственном счете (аудитлог того, за что были списаны или начислены деньги, с подробным описанием каждой из задач + текущий баланс) (query, read model)*  
- При ассайне на попуга с него списывается стоимость таска (command)  
Actor   "Task.Created" event, "Tasks.Reassigned" event  
Command Assigne task  
Data    Task, Account  
Event   Account.MoneyWithdrawn  
- Запись в лог операций по балансу  
Actor   "Task.Created" event, "Tasks.Reassigned" event  
Command Create accounting record (save withdrawal)  
Data    Task, Account  
Event   Account.UpdateLog  
  
- После выполнения таска попуга начисляется рандомная сумма (command)  
Actor   "Task.Completed" event  
Command Complete task  
Data    Task, Account  
Event   Account.MoneyCredited  
- Запись в лог операций по балансу  
Actor   "Task.Completed" event  
Command Create accounting record (save credit)  
Data    Task, Account  
Event   Account.UpdateLog  
- В конце рабочего дня выполняется "снятие всех тасков"?  
Actor   Crontab  
Command Finish work  
Data    Account, Tasks  
Event   Work.Finished  
- Каждый сотрудник в конце дня должен получить инфо-письмо с суммой выплаты за прошедший день (command)  
Actor   "Work.Finished" event  
Command Send email with day payments  
Data    Account, Accounting  
Event   Notification.EmailSended  
- Каждый сотрудник в конце дня получает обновление в аудитлог и обнуление баланса (кроме отрицательного баланса, его мы переносим в завтрашний день) (command)  
Actor   "Work.Finished" event  
Command Update auditlog and balance nullify  
Data    Account, Accounting  
Event   Account.AuditlogReseted  
- *Аналитику могут посмотреть только админы-попуги (query, read model)*  
  
## 3. Набрасываем модель данных  
- Account, Account role, Account balance, Account auth info  
- Task, Task description, Task state, Task price, Task assigned account  
- Accounting (Account auditlog)  
- Notification (Accounting auditlog, Account email)  
  
## 4. Выделяем домены (по акторам/контексту)  
- Auth domain  
    - Account  
        - id  
        - login  
        - pass  
        - roles  
        - balance  
- Task domain  
    - [Account - another domain piece]  
        - id  
        - role  
    - Task  
        - id  
        - assigned account id  
        - description  
        - state  
        - price (withdrawal cost)  
    - Reassigne  
        - Task (id) list   
        - Account (id) list (not manager, accountant, admin roles)  
    - **Money withdrawal event**    
        - Account id  
        - Task state - assigned (для простоты просто ассайн, не будем указывать причину ассайна - создание таска или реассайн)  
        - Task price  
        - DateTime  
    - **Money credited event**  
        - Account id  
        - Task state - completed  
        - ~~Task price~~ rand(20..40)$  
        - DateTime  
- Ишддштп (Accounting) domain  
    - [Account - another domain piece]  
        - id  
        - roles  
    - Accounting (auditlog)  
        - Account id  
        - Task id  
        - Task state (from Money withdrawal/credited event, пока пусть будет дублирование инфы - тип события и статус. В будущем можно это поле не делать)  
        - Price (from Money credited event?)  
        - DateTime (from event!)  
    - Work finish (crontab)  
        - Account (id, email) list (not manager, accountant, admin roles)  
        - Accounting (auditlog) list  
    - **Account today balance calculated event**  
        - Account id  
        - Auditlog  
    - Update auditlog and balance nullify  
        - Account id  
        - Auditlog (добавляем запись что сумма была выплачена)  
- Notification domain  
    - [Account - another domain piece]  
        - id  
        - roles  
        - email  
    - Notify  
        - Account id  
        - Auditlog  
  
## 5. Примеряемся как разделить домены на сервисы  
- Auth
- Task
- Accounting
- Notification

## 6. Определяем коммуникации между сервисами  
Выбираем **асинхронный** подход с отправкой событий в очередь - сервисы сами будут что-то делать, когда появится событие.    
Обновление данных учетных записей в в сервисах тоже **асинхронное**, считаем что нам не важна небольшая возможная задержка.    
Данные для коммуникаци между сервисами описаны в п.4 "Выделяем домены"  
  