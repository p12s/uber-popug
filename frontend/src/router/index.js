import Tasks from "../pages/Tasks";
import Login  from "../pages/Login";
import TaskDetail from '../pages/TaskDetail';

export const privateRoutes = [
  {path: '/tasks', component: Tasks, exact: true},
  {path: '/tasks/:id', component: TaskDetail, exact: true},
]

export const publicRoutes = [
  {path: '/login', component: Login, exact: true},
]