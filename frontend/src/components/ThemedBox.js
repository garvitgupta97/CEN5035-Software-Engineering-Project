import { Box, useColorMode } from '@chakra-ui/react';
import darkbc from '.././dark.png'
import sky from '.././skylight.png'

const ThemedBox = ({ light = 'blue.500', dark = 'gray.500', children, ...rest }) => {
  const { colorMode } = useColorMode();
  return (
    <Box
    backgroundImage = {colorMode === 'light' ? `url(${sky})` : `url(${darkbc})`}
    backgroundColor={ colorMode === 'light' ? light : dark}

     
      {...rest}
    >
      {children}
    </Box>
  );
}

export default ThemedBox;