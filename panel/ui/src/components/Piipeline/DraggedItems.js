import { CloseLarge } from '@carbon/icons-react';
import { TextInput } from '@carbon/react';
import React from 'react';

const DraggedItems = ({
  usedProcessors,
  onDragStart,
  onDragEnter,
  onDragEnd,
  setUsedProcessors,
}) => {
  return (
    <>
      {Boolean(usedProcessors?.length) &&
        usedProcessors.map((item, index) => {
          return (
            <div
              key={item?.id}
              onDragStart={(event) =>
                onDragStart({
                  id: item?.id,
                  event,
                  position: index,
                  source: 'used_processors',
                })
              }
              onDragOver={(e) => e.preventDefault()}
              onDragEnter={(event) =>
                onDragEnter({
                  id: item?.id,
                  event,
                  position: index,
                  source: 'used_processors',
                })
              }
              onDragEnd={(event) => onDragEnd({ id: item?.id, event })}
              draggable
            >
              <div
                style={{
                  maxWidth: '400px',
                  padding: '20px 20px 20px 20px',
                  borderTopRightRadius: '10px',
                  borderTopLeftRadius: '10px',
                  background: '#000',
                  display: 'flex',
                  justifyItems: 'center',
                  justifyContent: 'space-between',
                  alignItems: 'center',
                }}
              >
                <p
                  style={{
                    fontSize: '16px',
                    fontWeight: 'bold',
                    color: 'white',
                  }}
                >
                  {item.name}
                </p>
                <CloseLarge
                  size={24}
                  style={{ color: 'white', padding: '2px' }}
                  onClick={(e) => {
                    setUsedProcessors(
                      usedProcessors.filter((value, index1) => {
                        return index1 !== index;
                      })
                    );
                  }}
                />
              </div>
              {item?.arguments?.length ? (
                <div
                  style={{
                    maxWidth: '400px',
                    borderBottomRightRadius: '10px',
                    borderBottomLeftRadius: '10px',
                    paddingTop: '20px',
                    paddingBottom: '10px',
                    background: '#f4f4f4',
                    marginBottom: '30px',
                  }}
                >
                  {item?.arguments
                    ? item?.arguments.map?.((arg, argIndex) => (
                        <div
                          key={arg.name}
                          style={{
                            marginBottom: '20px',
                            marginLeft: '20px',
                            marginRight: '20px',
                          }}
                        >
                          {arg.type === 'Text' ? (
                            <TextInput
                              type="text"
                              onChange={(e) => {
                                arg.value = e.target.value;
                                item.arguments[argIndex] = arg;
                                usedProcessors[index] = item;
                                setUsedProcessors(
                                  (prevState) => usedProcessors
                                );
                              }}
                              required={arg.required}
                              placeholder={arg.name}
                              defaultValue={arg.value}
                            />
                          ) : arg.type === 'Numeric' ? (
                            <TextInput
                              type="text"
                              onChange={(e) => {
                                arg.value = parseInt(e.target.value);
                                item.arguments[argIndex] = arg;
                                usedProcessors[index] = item;
                                setUsedProcessors(
                                  (prevState) => usedProcessors
                                );
                              }}
                              required={arg.required}
                              defaultValue={arg.value}
                              placeholder={arg.name + ' (Number)'}
                            />
                          ) : (
                            <p>No input found for argument {arg.name}</p>
                          )}

                          {/*{arg.required ? <p style={{ paddingTop: '3px', color: 'red' }}>*</p> : null}*/}
                          {/*<p style={{ fontWeight: 'bold', padding: '3px 10px 0 10px' }}>{arg.type}</p>*/}
                          {/*<p*/}
                          {/*  style={{ paddingTop: '3px' }}>{arg.name !== 'processable.body' ? arg.name : ''}</p>*/}
                        </div>
                      ))
                    : null}
                </div>
              ) : (
                <p>No arguments</p>
              )}
            </div>
          );
        })}{' '}
    </>
  );
};

export default DraggedItems;
