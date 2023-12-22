import { Modal, TextInput } from '@carbon/react';
import React, { forwardRef, useImperativeHandle, useState } from 'react';

const PipelinePopup = forwardRef(({ pipelineName, setPipelineName }, ref) => {
  const [openPipeline, setOpenPipeline] = useState(true);

  useImperativeHandle(
    ref,
    () => {
      return { pipelineName };
    },
    []
  );

  return (
    <Modal
      open={openPipeline}
      preventCloseOnClickOutside={true}
      isFullWidth
      modalHeading="Create a new pipeline"
      modalLabel="Pipeline information"
      primaryButtonText="Continue"
      secondaryButtonText="Back to pipelines"
      onRequestSubmit={(event) => {
        if (pipelineName) {
          setOpenPipeline(false);
        }
      }}
      onRequestClose={(e) => {
        window.location.href = '/pipelines';
      }}>
      <div
        style={{
          padding: '20px',
        }}>
        <TextInput
          required={true}
          data-modal-primary-focus
          id="text-input-1"
          labelText="Pipeline name"
          placeholder="e.g. LOGIN_USER__V1"
          value={pipelineName}
          style={{
            marginBottom: '1rem',
          }}
          onChange={(event) => setPipelineName(event.target.value)}
        />
      </div>
    </Modal>
  );
});

export default PipelinePopup;
