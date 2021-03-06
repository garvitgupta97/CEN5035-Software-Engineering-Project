import React from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';
import {
  Box,
  Stack,
  FormControl,
  FormErrorMessage,
  Input,
  Textarea,
  RadioGroup,
  Radio,
  Select,
  Button,
  Alert,
  AlertIcon,
} from '@chakra-ui/react';
import {
  createLoadingAndErrorSelector
} from '../selectors';
import { submitPost } from '../actions/post';

class CreatePostPage extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      postType: 'text',
      title: '',
      body: '',
      url: ''
    };
  }

 

  handleSubmit = async (e) => {
    try {
      e.preventDefault();
      const { postType, title, body, url,subreddit } = this.state;
      const { submitPost, history } = this.props;
      const { id } = await submitPost({
        type: postType,
        title,
        body: postType === 'text' ? body : url,
        subreddit,
      });
      history.push(`/`);
    } catch (err) {}
  };

  render() {
    const {title, body} = this.state;
    const {
      srIsLoading,
      
      submitIsLoading,
      submitError,
      
    } = this.props;
    return (
      <Box w={['100%', '90%', '80%', '70%']} m="auto">
        {submitError && (
          <Alert status="error" mb={4}>
            <AlertIcon />
            {submitError}
          </Alert>
        )}
        <form onSubmit={this.handleSubmit}>
          <Stack spacing={3}>
            <FormControl>
              <Input
                value={title}
                onChange={(e) => this.setState({ title: e.target.value })}
                type="text"
                variant="filled"
                placeholder="title"
                isRequired
              />
            </FormControl>
            <FormControl>
                <Textarea
                  value={body}
                  onChange={(e) => this.setState({ body: e.target.value })}
                  variant="filled"
                  placeholder="text"
                  rows={10}
                  isRequired
                />
            </FormControl>
            
            <Button
              type="submit"
              isLoading={srIsLoading || submitIsLoading || null}
            >
              Submit
            </Button>
          </Stack>
        </form>
      </Box>
    );
  }
}

const {
  loadingSelector: srLoadingSelector,
  errorSelector: srErrorSelector,
} = createLoadingAndErrorSelector(['GET_SUBREDDITS']);

const {
  loadingSelector: submitLoadingSelector,
  errorSelector: submitErrorSelector,
} = createLoadingAndErrorSelector(['SUBMIT_POST'], false);

const mapStateToProps = (state) => ({
  srIsLoading: srLoadingSelector(state),
  srError: srErrorSelector(state),
  submitIsLoading: submitLoadingSelector(state),
  submitError: submitErrorSelector(state)
});

const mapDispatchToProps = (dispatch) => ({
  
  submitPost: (postDetails) => dispatch(submitPost(postDetails)),
});

export default withRouter(
  connect(mapStateToProps, mapDispatchToProps)(CreatePostPage)
);
