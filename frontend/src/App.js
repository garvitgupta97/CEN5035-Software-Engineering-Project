import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import { ChakraProvider, Box, Flex } from '@chakra-ui/react';
import theme from './theme';
import PublicRoute from './components/PublicRoute';
import PrivateRoute from './components/PrivateRoute';
import Navbar from './components/Navbar';
import CommentsPage from './components/CommentsPage';
import LoginPage from './components/LoginPage';
import RegisterPage from './components/RegisterPage';
import CreatePostPage from './components/CreatePostPage';
import PostList from './components/PostList';
import ThemedBox from './components/ThemedBox';

function App() {
  return (
    
    <ChakraProvider theme={theme}>
      <Router>
      <Navbar />
        <ThemedBox minHeight="100vh" >
       
          <Flex justifyContent="center">
            <Box width={['95%', '80%', '70%', '60%']} mb={10}>
              <Switch>
                <Route path="/comments/:id">
                  <CommentsPage />
                </Route>
                <PublicRoute path="/login">
                  <LoginPage />
                </PublicRoute>
                <PublicRoute path="/register">
                  <RegisterPage />
                </PublicRoute>
                <PrivateRoute path="/submit">
                  <CreatePostPage />
                </PrivateRoute>
                <Route path="/">
                  <PostList />
                </Route>
              </Switch>
            </Box>
          </Flex>
        </ThemedBox>
      </Router>
    </ChakraProvider>
  );
}

export default App;
