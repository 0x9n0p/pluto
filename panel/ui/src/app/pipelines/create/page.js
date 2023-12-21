'use client';

import { Content, Grid, Theme } from '@carbon/react';
import MainHeader from '@/components/MainHeader/MainHeader';
import React, { useRef, useState } from 'react';

import { DragDropContext } from 'react-beautiful-dnd';

import PipelinePopup from '../../../components/CreatePipeline/PipelinePopup';
import PiplineNavbar from '../../../components/CreatePipeline/Navbar';
import Processors from '../../../components/CreatePipeline/Processors';
import UsedProcessors from '../../../components/CreatePipeline/UsedProcessors';
import uuidv4 from '../../../utils/uuidv4';

export default function CreatePipelinePage() {
  const ProcessorsRef = useRef();
  const UsedProcessorsRef = useRef();
  const [pipelineName, setPipelineName] = useState('');
  const [errorMessageForSavePipeline, setErrorMessageForSavePipeline] =
    useState('');

  if (typeof window !== 'undefined')
    if (!localStorage.getItem('token')) window.location.assign('/auth');

  const onDragEnd = (result) => {
    const setUsedProcessors = UsedProcessorsRef.current.setUsedProcessors;
    const usedProcessors = [...UsedProcessorsRef.current.usedProcessors];

    const { destination, source, draggableId, type } = result;

    if (!destination) return;
    debugger;
    if (source?.droppableId === 'used_processors') {
      const newAdded = usedProcessors.find((item) => item?.id === draggableId);
      const newItems = new Array(...usedProcessors);
      newItems.splice(source.index, 1);
      newItems.splice(destination.index, 0, newAdded);
      setUsedProcessors(newItems);
      return;
    }

    if (destination?.droppableId === 'used_processors') {
      const newAdded = JSON.parse(result.draggableId);
      newAdded['id'] = uuidv4();
      if (newAdded?.arguments?.length) {
        newAdded.arguments.forEach((item) => {
          item['id'] = uuidv4();
        });
      }
      if (newAdded?.output?.length) {
        newAdded.output.forEach((item) => {
          item['id'] = uuidv4();
        });
      }
      if (newAdded?.input?.length) {
        newAdded.input.forEach((item) => {
          item['id'] = uuidv4();
        });
      }
      const newItems = new Array(...usedProcessors);
      newItems.splice(destination.index, 0, { ...newAdded });
      setUsedProcessors(newItems);
      return;
    }
  };
  return (
    <>
      <div>
        <Theme theme="g100">
          <MainHeader />
        </Theme>
        <Content>
          <PipelinePopup
            pipelineName={pipelineName}
            setPipelineName={setPipelineName}
          />

          <Grid className="create-page" fullWidth>
            <PiplineNavbar
              UsedProcessorsRef={UsedProcessorsRef}
              pipelineName={pipelineName}
              setErrorMessageForSavePipeline={setErrorMessageForSavePipeline}
            />
            <DragDropContext onDragEnd={onDragEnd}>
              {/* used - processors */}
              <UsedProcessors
                ref={UsedProcessorsRef}
                setErrorMessageForSavePipeline={setErrorMessageForSavePipeline}
                errorMessageForSavePipeline={errorMessageForSavePipeline}
              />

              {/* processors */}

              <Processors ref={ProcessorsRef} />
            </DragDropContext>
          </Grid>
        </Content>
      </div>
    </>
  );
}
