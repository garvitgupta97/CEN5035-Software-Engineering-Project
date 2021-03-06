import { useEffect } from 'react';
import { connect } from 'react-redux';
import { useParams } from 'react-router-dom';
import {
  Box,
  Flex,
  Text,
  Heading,
  Alert,
  AlertIcon,
  CircularProgress,
} from '@chakra-ui/react';
import Post from './Post';
import CommentsThread from './CommentsThread';
import WriteCommentBox from './WriteCommentBox';
import LoginAndRegisterButtons from './LoginAndRegisterButtons';
import {
  createLoadingAndErrorSelector,
  postSelector,
  commentsSelector,
  userSelector,
} from '../selectors';
import { getPostAndComments } from '../actions';

const getCommentsWithChildren = (comments) => {
  const commentsWithChildren = comments.map((comment) => ({
    ...comment,
    children: [],
  }));

  commentsWithChildren.forEach((childComment) => {
    const { parent_comment_id } = childComment;
    if (parent_comment_id && parent_comment_id != -1) {
      const parent = commentsWithChildren.find(
        (comment) => parent_comment_id === comment.id
      );
      parent.children = parent.children.concat(childComment);
    }
  });
  return commentsWithChildren.filter(
    ({ parent_comment_id, body, children }) =>
      parent_comment_id == -1 && (body !== null || children.length > 0)
  );
};

const CommentsPage = ({
  isLoading,
  error,
  post,
  comments,
  getPostAndComments,
  user,
}) => {
  const { id } = useParams();
  useEffect(() => {
    getPostAndComments(id);
  }, [getPostAndComments, id]);

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
  const {
    PostId,
    UserId,
    UserEmail,
    CreatedAt,
    Title,
    Content,
    Votes,
    has_voted,
  } = post;
  const numComments = comments.filter(({ body }) => body !== null).length;

  const rootComments = getCommentsWithChildren(comments);
  return (
    <Box>
      <Post
        id={PostId}
        type={'text'}
        author={UserEmail}
        createdAt={CreatedAt}
        title={Title}
        body={Content}
        numComments={numComments}
        numVotes={Votes}
        hasVoted={has_voted}
      />
      <br />
      {user ? (
        <Box>
          <Box m={2}>
            <Text as="span" color="gray.500">
              {'Comment as '}
            </Text>
            <Text as="span">{user.username}</Text>
          </Box>
          <WriteCommentBox postId={PostId} parentCommentId={null} />
        </Box>
      ) : (
        <Flex
          p={5}
          border="1px"
          borderRadius={5}
          borderColor="gray.500"
          alignItems="center"
          justifyContent="space-between"
        >
          <Heading size="md" color="gray.500">
            Log in or register to leave a comment
          </Heading>
          <LoginAndRegisterButtons />
        </Flex>
      )}
      <br />
      <CommentsThread comments={rootComments} />
    </Box>
  );
};

const { loadingSelector, errorSelector } = createLoadingAndErrorSelector([
  'GET_POST_AND_COMMENTS',
]);

const mapStateToProps = (state) => ({
  isLoading: loadingSelector(state),
  error: errorSelector(state),
  post: postSelector(state),
  comments: commentsSelector(state),
  user: userSelector(state),
});

const mapDispatchToProps = (dispatch) => ({
  getPostAndComments: (id) => dispatch(getPostAndComments(id)),
});

export default connect(mapStateToProps, mapDispatchToProps)(CommentsPage);
