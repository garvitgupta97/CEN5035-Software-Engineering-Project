import { useEffect } from 'react';
import { connect } from 'react-redux';
import { Box } from '@chakra-ui/react';
import { Link, useLocation } from 'react-router-dom';
import {
  Heading,
  Spacer,
  HStack,
  Button,
  Menu,
  MenuButton,
  MenuList,
  MenuItem
} from '@chakra-ui/react';
import { ChevronDownIcon } from '@chakra-ui/icons';
import { ColorModeSwitcher } from '../ColorModeSwitcher';
import ThemedBox from './ThemedBox';
import {
  userSelector,
  createLoadingAndErrorSelector,
} from '../selectors';
import { startLogout } from '../actions/auth';
import LoginAndRegisterButtons from './LoginAndRegisterButtons';
import {Logo} from '../Logo'

const Navbar = ({
  user,
  isLoading,
  error,
  startLogout
}) => {
  const location = useLocation();

  return (
    <Box
      py={2}
      px={[0, 0, 10, 10]}
      display="flex"
      justifyContent="flex-start"
      alignItems="center"
      
      // mb={7}
    >
        {/* <Heading
          ml={[2, 4]}
          display={user ? 'block' : ['none', 'block']}
          fontSize={['1.3rem', '2.25rem']}
          mr={[2, 4]}
        >
          StudentUniverse
        </Heading> */}
      <Logo  />
      <HStack>
      
        <Button display={['none', 'flex']} as={Link} to="/">
          Home
        </Button>

        {user && (
          <Button display={['none', 'flex']} as={Link} to="/submit">
            Submit
          </Button>
        )}
      </HStack>
      <Spacer />

      {user ? (
        <Menu>
          <MenuButton as={Button} rightIcon={<ChevronDownIcon />}>
            {user.username}
          </MenuButton>
          <MenuList>
            <MenuItem display={['block', 'none']} as={Link} to="/submit">
              Submit post
            </MenuItem>
            <MenuItem
              onClick={async () => {
                await startLogout();
              }}
            >
              Logout
            </MenuItem>
          </MenuList>
        </Menu>
      ) : (
        <LoginAndRegisterButtons />
      )}
      <ColorModeSwitcher />
    </Box>
  );
};

const { loadingSelector, errorSelector } = createLoadingAndErrorSelector([
  'GET_SUBREDDITS',
]);

const mapStateToProps = (state) => ({
  isLoading: loadingSelector(state),
  error: errorSelector(state),
  user: userSelector(state),
});

const mapDispatchToProps = (dispatch) => ({
  startLogout: () => dispatch(startLogout())
});

export default connect(mapStateToProps, mapDispatchToProps)(Navbar);
