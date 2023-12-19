import {
  Breadcrumb,
  BreadcrumbItem,
  Button,
  Column,
  Grid,
} from '@carbon/react';
import axios from 'axios';
import React from 'react';
import { Address } from '../../../settings';

const PiplineNavbar = ({
  pipelineName,
  usedProcessors,
  UsedProcessorsRef,
  setErrorMessageForSavePipeline,
}) => {
  return (
    <Column
      lg={16}
      md={8}
      sm={4}
      className="create-page_header"
      style={{ marginBottom: '48px' }}
    >
      <Breadcrumb>
        <BreadcrumbItem>
          <a href="/">Home</a>
        </BreadcrumbItem>
        <BreadcrumbItem>
          <a href="/pipelines">Pipelines</a>
        </BreadcrumbItem>
      </Breadcrumb>
      <Grid fullWidth>
        <Column md={4} lg={{ span: 7, offset: 0 }} sm={4}>
          <h1 className="create-page__heading">Create a new pipeline</h1>
        </Column>
        <Column md={4} lg={{ span: 1, offset: 13 }} sm={4}>
          <Button
            onClick={(event) => {
              axios
                .post(
                  Address() + '/api/v1/pipelines',
                  {
                    name: pipelineName,
                    processors: UsedProcessorsRef?.current?.usedProcessors,
                  },
                  {
                    headers: {
                      Authorization: `Bearer ${localStorage.getItem('token')}`,
                    },
                  }
                )
                .then(function (response) {
                  if (response.status !== 201) {
                    setErrorMessageForSavePipeline(
                      'Unexpected response from server'
                    );
                    return;
                  }
                  window.location.href = '/pipelines';
                })
                .catch(function (error) {
                  if (error.response) {
                    setErrorMessageForSavePipeline(error.response.data.message);
                  } else {
                    setErrorMessageForSavePipeline('Unknown Error');
                  }
                });
            }}
          >
            Save {pipelineName}
          </Button>
        </Column>
      </Grid>
    </Column>
  );
};

export default PiplineNavbar;
