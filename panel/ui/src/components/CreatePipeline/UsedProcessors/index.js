import { Column, InlineNotification } from '@carbon/react';
import React, { forwardRef } from 'react';
import { Droppable } from 'react-beautiful-dnd';
import UsedProcessorList from './UsedProcessorList';
import uuidv4 from '../../../utils/uuidv4';

const UsedProcessors = forwardRef(
  ({ setErrorMessageForSavePipeline, errorMessageForSavePipeline }, ref) => {
    return (
      <Droppable
        droppableId={'used_processors'}
        direction="vertical"
        mode="standard">
        {(provided) => (
          <Column
            md={{ span: 6, offset: 1 }}
            lg={{ span: 6, offset: 1 }}
            sm={{ span: 6 }}>
            {errorMessageForSavePipeline !== '' && (
              <InlineNotification
                aria-label="closes notification"
                kind="error"
                statusIconDescription="notification"
                subtitle={errorMessageForSavePipeline}
                onClose={() => {
                  setErrorMessageForSavePipeline('');
                }}
                style={{
                  marginBottom: '16px',
                  maxWidth: '500px',
                }}
              />
            )}

            <div
              {...provided.droppableProps}
              ref={provided.innerRef}
              style={{
                paddingBottom: '30px',
                paddingTop: '30px',
              }}>
              <UsedProcessorList
                ref={ref}
                setErrorMessageForSavePipeline={setErrorMessageForSavePipeline}
                errorMessageForSavePipeline={errorMessageForSavePipeline}
              />
              {provided.placeholder}
            </div>
          </Column>
        )}
      </Droppable>
    );
  }
);

export default UsedProcessors;
