import React from 'react';
import { Image, keyframes, usePrefersReducedMotion } from '@chakra-ui/react';
import logo from './logo3.png';



export const Logo = props => {


  return <Image src={logo} 
  width="50"
  height="50"
  marginRight={2}
  className="d-inline-block align-center"
  {...props} />;
};
