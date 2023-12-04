'use client';

import {
  Breadcrumb,
  BreadcrumbItem,
  Grid,
  Column,
  Search,
  ContainedList,
  ContainedListItem,
  Layer,
} from '@carbon/react';
import { useEffect, useState } from 'react';

export default function PipelineCreatorPage() {
  const [searchTerm, setSearchTerm] = useState('');
  const [searchResults, setSearchResults] = useState([]);
  const handleChange = (event) => {
    setSearchTerm(event.target.value);
  };

  const listItems = [
    'List item 1',
    'List item 2',
    'List item 3',
    'List item 4',
  ];

  useEffect(() => {
    const results = listItems.filter((listItem) =>
      listItem.toLowerCase().includes(searchTerm.toLowerCase()),
    );
    setSearchResults(results);
  }, [searchTerm]);

  return (
    <Grid className='create-page' fullWidth>

      <Column lg={16} md={8} sm={4} className='create-page_header'>
        <Breadcrumb noTrailingSlash>
          <BreadcrumbItem>
            <a href='/'>Home</a>
          </BreadcrumbItem>
          <BreadcrumbItem>
            <a href='/pipelines'>Pipelines</a>
          </BreadcrumbItem>
          <BreadcrumbItem>
            Create a new pipeline
          </BreadcrumbItem>
        </Breadcrumb>
        <h1 className='create-page__heading'>Create a new pipeline</h1>
      </Column>

      <Column lg={16} md={8} sm={4} className='create-page_main'>
        <Grid>

          <Column md={4} lg={7} sm={4}>
            {/*<ContainedList kind='on-page' action={''}>*/}
            {/*  {searchResults.map((listItem, key) => (*/}
            {/*    <ContainedListItem key={key}>{listItem}</ContainedListItem>*/}
            {/*  ))}*/}
            {/*</ContainedList>*/}
            <div
              style={
                {
                  marginTop: '48px',
                  height: '100px',
                  border: '2px dashed gray',
                  borderColor: '#d2d2d2',
                  borderRadius: '5px',
                  display: 'flex',
                  alignItems: 'center',
                  justifyContent: 'center',
                }
              }
            >
              <p>Drop a processor here</p>
            </div>
          </Column>

          <Column md={4} lg={{ span: 6, offset: 8 }} sm={4}>
            <Layer as='section'>
              <ContainedList label='Processors' kind='on-page' action={''}>
                <Search
                  placeholder={`Search between ${listItems.length} processors`}
                  value={searchTerm}
                  onChange={handleChange}
                  closeButtonLabelText='Clear search input'
                  size='lg'
                  labelText='Processors'
                />
                {searchResults.map((listItem, key) => (
                  <ContainedListItem key={key}>{listItem}</ContainedListItem>
                ))}
              </ContainedList>
            </Layer>
          </Column>

        </Grid>
      </Column>

    </Grid>
  );
};

