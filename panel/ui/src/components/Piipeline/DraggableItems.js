import { categoryToColor } from '@/app/pipelines/create/page';
import {
  ExpandableTile,
  Tag,
  TileAboveTheFoldContent,
  TileBelowTheFoldContent,
} from '@carbon/react';

const DraggableItems = ({
  searchResults,
  onDragStart,
  onDragEnter,
  handleDragEnd,
}) => {
  return (
    <>
      {Boolean(searchResults?.length) &&
        searchResults.map((item, index) => {
          return (
            <div
              key={item?.id}
              onDragStart={(event) => {
                onDragStart({
                  id: item?.id,
                  event,
                  position: index,
                  source: 'processors',
                });
                // handleDragStart(e, index, 'processors');
              }}
              onDragOver={(e) => e.preventDefault()}
              onDragEnter={(event) =>
                onDragEnter({
                  id: item?.id,
                  event,
                  position: index,
                  source: 'processors',
                })
              }
              onDragEnd={(event) => handleDragEnd({ event, id: item?.id })}
              draggable
            >
              <ExpandableTile
                style={{
                  paddingLeft: '20px',
                  marginTop: '10px',
                  marginBottom: '10px',
                }}
                tileCollapsedIconText="Details"
                tileExpandedIconText="Details"
              >
                <TileAboveTheFoldContent>
                  <div>
                    <h5>{item.name}</h5>
                    <p
                      style={{
                        fontSize: '14px',
                      }}
                    >
                      {item.description}
                    </p>
                    <Tag type={categoryToColor(item.category)}>
                      {item.category}
                    </Tag>
                  </div>
                </TileAboveTheFoldContent>
                <TileBelowTheFoldContent>
                  <div style={{ marginTop: '10px' }}>
                    <p
                      style={{
                        fontWeight: 'bold',
                        fontSize: '18px',
                      }}
                    >
                      Arguments
                    </p>
                    {item.arguments ? (
                      item.arguments.map((arg, index) => (
                        <div key={arg.name} style={{ display: 'flex' }}>
                          {arg.required ? (
                            <p
                              style={{
                                paddingTop: '3px',
                                color: 'red',
                              }}
                            >
                              *
                            </p>
                          ) : null}
                          <p
                            style={{
                              fontWeight: 'bold',
                              padding: '3px 10px 0 10px',
                            }}
                          >
                            {arg.type}
                          </p>
                          <p style={{ paddingTop: '3px' }}>
                            {arg.name !== 'processable.body' ? arg.name : ''}
                          </p>
                        </div>
                      ))
                    ) : (
                      <p>No arguments</p>
                    )}
                  </div>

                  <div style={{ marginTop: '10px' }}>
                    <p
                      style={{
                        fontWeight: 'bold',
                        fontSize: '18px',
                      }}
                    >
                      Input
                    </p>
                    {item.input ? (
                      item.input.map((arg, index) => (
                        <div key={arg.name} style={{ display: 'flex' }}>
                          {arg.required ? (
                            <p
                              style={{
                                paddingTop: '3px',
                                color: 'red',
                              }}
                            >
                              *
                            </p>
                          ) : null}
                          <p
                            style={{
                              fontWeight: 'bold',
                              padding: '3px 10px 0 10px',
                            }}
                          >
                            {arg.type}
                          </p>
                          <p style={{ paddingTop: '3px' }}>
                            {arg.name !== 'processable.body' ? arg.name : ''}
                          </p>
                        </div>
                      ))
                    ) : (
                      <p>No input</p>
                    )}
                  </div>

                  <div style={{ marginTop: '10px' }}>
                    <p
                      style={{
                        fontWeight: 'bold',
                        fontSize: '18px',
                      }}
                    >
                      Output
                    </p>
                    {item.output ? (
                      item.output.map((arg, index) => (
                        <div key={arg.name} style={{ display: 'flex' }}>
                          {arg.required ? (
                            <p
                              style={{
                                paddingTop: '3px',
                                color: 'red',
                              }}
                            >
                              *
                            </p>
                          ) : null}
                          <p
                            style={{
                              fontWeight: 'bold',
                              padding: '3px 10px 0 10px',
                            }}
                          >
                            {arg.type}
                          </p>
                          <p style={{ paddingTop: '3px' }}>
                            {arg.name !== 'processable.body' ? arg.name : ''}
                          </p>
                        </div>
                      ))
                    ) : (
                      <p>No output</p>
                    )}
                  </div>
                </TileBelowTheFoldContent>
              </ExpandableTile>
            </div>
          );
        })}
    </>
  );
};

export default DraggableItems;
