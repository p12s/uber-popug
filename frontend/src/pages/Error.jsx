import React from 'react';
import { Empty } from '@douyinfe/semi-ui';
import { IllustrationNotFound, IllustrationNotFoundDark } from '@douyinfe/semi-illustrations';

const Error = () => {
  return (
    <div>
      <h1>404. Not found</h1>
      <Empty
        image={<IllustrationNotFound style={{ width: 150, height: 150 }} />}
        darkModeImage={<IllustrationNotFoundDark style={{ width: 150, height: 150 }} />}
        description={'Page 404'}
        style={emptyStyle} />
    </div>
  );
}

export default Error;
