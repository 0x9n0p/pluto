import {
  Column,
  ContainedList,
  InlineNotification,
  Search,
} from '@carbon/react';
import React, { forwardRef, useState } from 'react';
import { Droppable } from 'react-beautiful-dnd';
import StickyBox from 'react-sticky-box';
import ProcessorList from './ProcessorList';

const Processors = forwardRef((_, ref) => {
  const [errorMessage, setErrorMessage] = useState('');
  const [searchTerm, setSearchTerm] = useState('');

  return (
    <Droppable droppableId={'processors'} direction="vertical">
      {(provided) => (
        <Column md={4} lg={{ span: 6, offset: 8 }} sm={4}>
          <StickyBox
            style={{ height: '20px' }}
            offsetTop={100}
            offsetBottom={20}
          >
            <ContainedList label="Processors" kind="on-page" action={''}>
              <Search
                placeholder="Filter"
                closeButtonLabelText="Clear search input"
                size="lg"
                labelText="Filter search"
                value={searchTerm}
                onChange={(e) => {
                  setSearchTerm(e.target.value);
                }}
              />
              {errorMessage !== '' && (
                <InlineNotification
                  aria-label="closes notification"
                  kind="error"
                  statusIconDescription="notification"
                  subtitle={errorMessage}
                  onClose={() => {
                    setErrorMessage('');
                  }}
                  style={{ marginBottom: '16px' }}
                />
              )}

              <div {...provided.droppableProps} ref={provided.innerRef}>
                <ProcessorList
                  ref={ref}
                  setErrorMessage={setErrorMessage}
                  searchTerm={searchTerm}
                  setSearchTerm={setSearchTerm}
                />
                {provided.placeholder}
              </div>
            </ContainedList>
          </StickyBox>
        </Column>
      )}
    </Droppable>
  );
});

export default Processors;
