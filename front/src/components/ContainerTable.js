import React, { useEffect, useState } from 'react';
import { Table, Button, Spin, message } from 'antd';
import { getContainers } from '../services/api';

const ContainerTable = () => {
  const [containers, setContainers] = useState([]);
  const [loading, setLoading] = useState(false);
  

  // Функция для загрузки данных
  const fetchData = async () => {
    setLoading(true);
    try {
      const data = await getContainers();
      setContainers(data);
      message.success('Data loaded successfully!');
    } catch (error) {
      console.error('Error fetching containers:', error);
      message.error('Failed to load data.');
    } finally {
      setLoading(false);
    }
  };

  // Загружаем данные при монтировании компонента
  useEffect(() => {
    fetchData();
  }, []);

  // Колонки для таблицы
  const columns = [
    {
      title: 'IP Address',
      dataIndex: 'ip_address',
      key: 'ip_address',
    },
    {
      title: 'Ping Time',
      dataIndex: 'ping_time',
      key: 'ping_time',
    },
    {
      title: 'Last Successful Ping',
      dataIndex: 'last_successful_ping',
      key: 'last_successful_ping',
    },
  ];

  const tableStyle = {
    boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)',
    borderRadius: '8px',
    overflow: 'hidden',
  };

  return (
    <div style={containerStyle}>
    <h1 style={{ textAlign: 'center', marginBottom: '20px' }}>Docker Container Monitoring</h1>
    <Button
      type="primary"
      onClick={fetchData}
      loading={loading}
      style={{ marginBottom: '20px' }}
    >
      Refresh Data
    </Button>
    <Spin spinning={loading}>
    <Table
  dataSource={containers}
  columns={columns}
  rowKey="id"
  pagination={false}
  bordered
  style={tableStyle}
/>
    </Spin>
  </div>
  
  );
};

export default ContainerTable;