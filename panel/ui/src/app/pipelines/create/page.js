'use client';

import React, { Component } from 'react';
import { DragDropContext, Droppable, Draggable } from 'react-beautiful-dnd';
import { Breadcrumb, BreadcrumbItem, Column, Grid } from '@carbon/react';
import StickyBox from 'react-sticky-box';

// fake data generator
const getItems = (count, offset = 0) =>
  Array.from({ length: count }, (v, k) => k).map((k) => ({
    id: `item-${k + offset}`,
    content: `item ${k + offset}`,
  }));

// a little function to help us with reordering the result
const reorder = (list, startIndex, endIndex) => {
  const result = Array.from(list);
  const [removed] = result.splice(startIndex, 1);
  result.splice(endIndex, 0, removed);

  return result;
};

/**
 * Moves an item from one list to another list.
 */
const move = (source, destination, droppableSource, droppableDestination) => {
  const sourceClone = Array.from(source);
  const destClone = Array.from(destination);
  const [removed] = sourceClone.splice(droppableSource.index, 1);

  destClone.splice(droppableDestination.index, 0, removed);

  const result = {};
  result[droppableSource.droppableId] = sourceClone;
  result[droppableDestination.droppableId] = destClone;

  return result;
};

const grid = 8;

const getItemStyle = (isDragging, draggableStyle) => ({
  // some basic styles to make the items look a bit nicer
  userSelect: 'none',
  padding: grid * 2,
  margin: `0 0 ${grid}px 0`,

  // change background colour if dragging
  background: isDragging ? 'lightgreen' : 'grey',

  // styles we need to apply on draggables
  ...draggableStyle,
});

const getListStyle = (isDraggingOver) => ({
  background: isDraggingOver ? 'lightblue' : 'lightgrey',
  padding: grid,
  width: 250,
});

class CreatePipelinePage extends Component {
  state = {
    items: getItems(20),
    selected: getItems(15, 20),
  };

  /**
   * A semi-generic way to handle multiple lists. Matches
   * the IDs of the droppable container to the names of the
   * source arrays stored in the state.
   */
  id2List = {
    droppable: 'items',
    droppable2: 'selected',
  };

  getList = (id) => this.state[this.id2List[id]];

  onDragEnd = (result) => {
    const { source, destination } = result;

    // dropped outside the list
    if (!destination) {
      return;
    }

    if (source.droppableId === destination.droppableId) {
      const items = reorder(
        this.getList(source.droppableId),
        source.index,
        destination.index
      );

      let state = { items };

      if (source.droppableId === 'droppable2') {
        state = { selected: items };
      }

      this.setState(state);
    } else {
      const result = move(
        this.getList(source.droppableId),
        this.getList(destination.droppableId),
        source,
        destination
      );

      this.setState({
        items: result.droppable,
        selected: result.droppable2,
      });
    }
  };

  // Normally you would want to split things out into separate components.
  // But in this example everything is just done in one place for simplicity
  render() {
    return (
      <>
        {localStorage.getItem('token') ? (
          <Grid className="create-page" fullWidth>
            <Column
              lg={16}
              md={8}
              sm={4}
              className="create-page_header"
              style={{ marginBottom: '48px' }}
            >
              <Breadcrumb noTrailingSlash>
                <BreadcrumbItem>
                  <a href="/">Home</a>
                </BreadcrumbItem>
                <BreadcrumbItem>
                  <a href="/pipelines">Pipelines</a>
                </BreadcrumbItem>
                <BreadcrumbItem>Create a new pipeline</BreadcrumbItem>
              </Breadcrumb>
              <h1 className="create-page__heading">Create a new pipeline</h1>
            </Column>

            <DragDropContext onDragEnd={this.onDragEnd}>
              <Column md={4} lg={7} sm={4}>
                <Droppable droppableId="droppable">
                  {(provided, snapshot) => (
                    <div
                      ref={provided.innerRef}
                      style={getListStyle(snapshot.isDraggingOver)}
                    >
                      {this.state.items.map((item, index) => (
                        <Draggable
                          key={item.id}
                          draggableId={item.id}
                          index={index}
                        >
                          {(provided, snapshot) => (
                            <div
                              ref={provided.innerRef}
                              {...provided.draggableProps}
                              {...provided.dragHandleProps}
                              style={getItemStyle(
                                snapshot.isDragging,
                                provided.draggableProps.style
                              )}
                            >
                              {item.content}
                            </div>
                          )}
                        </Draggable>
                      ))}
                      {provided.placeholder}
                    </div>
                  )}
                </Droppable>

                {/* TODO: It must be visible if the pipeline is empty. */}
                {/*<div*/}
                {/*  style={*/}
                {/*    {*/}
                {/*      height: '100px',*/}
                {/*      border: '2px dashed gray',*/}
                {/*      borderColor: '#d2d2d2',*/}
                {/*      borderRadius: '5px',*/}
                {/*      display: 'flex',*/}
                {/*      alignItems: 'center',*/}
                {/*      justifyContent: 'center',*/}
                {/*    }*/}
                {/*  }*/}
                {/*>*/}
                {/*  <p>Drop a processor here</p>*/}
                {/*</div>*/}
              </Column>

              <Column md={4} lg={{ span: 6, offset: 8 }} sm={4}>
                <StickyBox offsetTop={100} offsetBottom={20}>
                  <Droppable droppableId="droppable2">
                    {(provided, snapshot) => (
                      <div
                        ref={provided.innerRef}
                        style={getListStyle(snapshot.isDraggingOver)}
                      >
                        {this.state.selected.map((item, index) => (
                          <Draggable
                            key={item.id}
                            draggableId={item.id}
                            index={index}
                          >
                            {(provided, snapshot) => (
                              <div
                                ref={provided.innerRef}
                                {...provided.draggableProps}
                                {...provided.dragHandleProps}
                                style={getItemStyle(
                                  snapshot.isDragging,
                                  provided.draggableProps.style
                                )}
                              >
                                {item.content}
                              </div>
                            )}
                          </Draggable>
                        ))}
                        {provided.placeholder}
                      </div>
                    )}
                  </Droppable>
                </StickyBox>
              </Column>
            </DragDropContext>
          </Grid>
        ) : (
          window.location.assign('/auth')
        )}
      </>
    );
  }
}

export default CreatePipelinePage;
