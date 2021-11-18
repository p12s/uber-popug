import React, {useContext} from 'react';
import { Redirect, Route, Switch } from 'react-router-dom';
import { AuthContext } from '../context';
import { publicRoutes, privateRoutes } from '../router';

const AppRouter = () => {
  const {isAuth} = useContext(AuthContext);

  return (
    isAuth
    ? <Switch>
        {privateRoutes.map(route => 
          <Route 
          key={route.path}
          component={route.component} 
          path={route.path} 
          exact={route.exact} />
        )}
        <Redirect to="/tasks"/>
      </Switch>
    : <Switch>
        {publicRoutes.map(route => 
          <Route 
          key={route.path}
          component={route.component} 
          path={route.path} 
          exact={route.exact} />
        )}
        <Redirect to="/login"/>
      </Switch>
  );
}

export default AppRouter;