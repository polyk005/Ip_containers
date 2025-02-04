import React from 'react';
import { ConfigProvider } from 'antd';
import ContainerTable from './components/ContainerTable';

const App = () => {
  return (
    <ConfigProvider
      theme={{
        token: {
          colorPrimary: '#1890ff', // Основной цвет (синий)
        },
      }}
    >
      <ContainerTable />
    </ConfigProvider>
  );
};

export default App;

const containerStyle = {
    maxWidth: '1200px',
    margin: '0 auto',
    padding: '20px',
  };