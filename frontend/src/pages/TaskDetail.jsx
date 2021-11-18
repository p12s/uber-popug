import React from 'react';
// import React, { useState, useEffect } from 'react';
// import {useParams} from 'react-router-dom';
// import TaskService from '../API/TaskService';
// import { useFetching } from '../hooks/useFetching';

const TaskDetail = () => {
  // const params = useParams();
  // const [post, setPost] = useState({ });
  // const [fetchPostById, isLoading, error] = useFetching( async (id) => {
  //   const response = await PostService.getById(id)
  //   setPost(response.data)
  // });

  // const [comments, setComments] = useState([]);
  // const [fetchComments, isCommentLoading, errorComment] = useFetching( async (id) => {
  //   const response = await PostService.getCommentsById(id)
  //   setComments(response.data)
  // });

  // useEffect(() => {
  //   fetchPostById(params.id)
  //   fetchComments(params.id)
  // }, []);

  return (
    <div>
      {/* <h1>PostIdPage {params.id}</h1> */}
      <h1>Task detail</h1>
    </div>
  );
}

export default TaskDetail;