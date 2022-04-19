import { useState } from 'react';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import moment from 'moment';
import {
  Text,
  Heading,
  Box,
  Flex,
  HStack,
  Tooltip,
  IconButton,
  useColorMode,
} from '@chakra-ui/react';
import { ChatIcon, EditIcon } from '@chakra-ui/icons';
import ThemedBox from './ThemedBox';
import UpvoteBar from './UpvoteBar';
import EditBox from './EditBox';
import DeleteButton from './DeleteButton';
import ChakraMarkdown from './ChakraMarkdown';
import { userSelector } from '../selectors';


const Post = ({
  id,
  type,
  author,
  createdAt,
  title,
  body,
  numVotes,
  hasVoted,
  numComments,
  user,
}) => {
  const { colorMode } = useColorMode();
  
  // post font color
  const postDetailColor = 'black';
  
  // button on post color
  const postDetailBgColor = colorMode === 'light' ? 'gray.200' : 'gray.600'; 
  

  const [isEditing, setIsEditing] = useState(false);
  const deletedText = '[deleted]';
  return (
    <ThemedBox
      p={4}
      borderRadius="md"
      width="100%"
      light="gray.100"
      dark="gray.400"
      backgroundImage="null"
    >
      <Flex>
        <UpvoteBar
          type="post"
          numVotes={numVotes}
          id={id}
          voteValue={hasVoted}
          color="black"
        />
        <Box flexGrow={1}>
          <Text as="span" color={postDetailColor}>
            {`Posted by `}
          </Text>
          <Text as="span">{author ? `${author}` : deletedText}</Text>
          <Text as="span" color={postDetailColor}>
            {' '}
            <Tooltip label={moment(createdAt).format('LLLL')}>
              {moment(createdAt).fromNow()}
            </Tooltip>
          </Text>
          <Heading
            display="block"
            href={body}
            mt={2}
            mb={4}
            fontSize="1.5em"
            fontWeight="500"
            color={postDetailColor}
          >
            {title || deletedText}
          </Heading>
          {isEditing ? (
              <EditBox
                
              type="post"
                id={id}
                initialText={body}
                onClose={() => setIsEditing(false)}
              />
            ) : (
              <Box listStylePosition="inside" color={postDetailColor}>
                <ChakraMarkdown >{body}</ChakraMarkdown>
              </Box>
            )
          }
          <Flex
            mt={3}
            alignItems="center"
            color={postDetailColor}
            fontWeight="bold"
          >
            <Box
              as={Link}
              to={`/comments/${id}`}
              p={2}
              borderRadius="sm"
              _hover={{ backgroundColor: postDetailBgColor }}
            >
              <ChatIcon mr={2} />
              {numComments} {numComments === 1 ? 'comment' : 'comments'}
            </Box>
          </Flex>
        </Box>
        {user && user.username === author && (
          <HStack alignItems="flex-start">
            {!isEditing && (
              <IconButton
                onClick={() => setIsEditing(true)}
                backgroundColor="inherit"
                
                icon={<EditIcon />}
              />
            )}
            <DeleteButton type="post" id={id} />
          </HStack>
        )}
      </Flex>
    </ThemedBox>
  );
};

const mapStateToProps = (state) => ({
  user: userSelector(state),
});

export default connect(mapStateToProps)(Post);
