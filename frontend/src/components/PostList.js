import React, { useEffect } from 'react';
import { connect } from 'react-redux';
import { useParams } from 'react-router-dom';
import {
  Box,
  Flex,
  Alert,
  AlertIcon,
  Heading,
  Text,
  CircularProgress,
} from '@chakra-ui/react';
import Post from './Post';
import { createLoadingAndErrorSelector, postListSelector } from '../selectors';
import { getPostList } from '../actions/postList';

const PostList = ({ isLoading, error, postList, getPostList }) => {

  useEffect(() => {
    getPostList({ });
  }, [getPostList]);

  if (isLoading) {
    return (
      <Flex m={10} justifyContent="center" alignItems="center">
        <CircularProgress isIndeterminate />
      </Flex>
    );
  } else if (error) {
    return (
      <Alert status="error">
        <AlertIcon />
        {error}
      </Alert>
    );
  }
  return (
    <Box>
      <Heading>{'Home'}</Heading>
      {postList.length > 0 ? (
        postList.map(
          ({
            PostId,
            Title,
            Content,
            Votes,
            hasVoted = 0,
            CommentsCount,
            Email,
            PostCreatedTime
          }) => (
            <Box key={`${PostId}-${Title}`} my={4}>
              <Post
                id={PostId}
                type={'text'}
                author={Email}
                createdAt={PostCreatedTime}
                title={Title}
                body={Content}
                numComments={CommentsCount}
                numVotes={Votes}
                hasVoted={hasVoted}
              />
            </Box>
          )
        )
      ) : (
        <Text m={5}>There are no posts to display.</Text>
      )}
    </Box>
  );
};

const { loadingSelector, errorSelector } = createLoadingAndErrorSelector([
  'GET_POST_LIST',
]);

const mapStateToProps = (state) => ({
  isLoading: loadingSelector(state),
  error: errorSelector(state),
  postList: postListSelector(state),
});

const mapDispatchToProps = (dispatch) => ({
  getPostList: (filters) => dispatch(getPostList(filters)),
});

export default connect(mapStateToProps, mapDispatchToProps)(PostList);
