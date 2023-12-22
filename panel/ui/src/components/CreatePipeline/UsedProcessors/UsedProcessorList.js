import { CloseLarge } from '@carbon/icons-react';
import { TextInput, Toggle } from '@carbon/react';
import React, { forwardRef, useImperativeHandle, useState } from 'react';
import { Draggable } from 'react-beautiful-dnd';
import uuidv4 from '../../../utils/uuidv4';
import { AnimatePresence, motion, stagger } from 'framer-motion';
import { varFade } from '../../animate/fade';

const UsedProcessorList = forwardRef((_, ref) => {
  const [usedProcessors, setUsedProcessors] = useState([]);

  useImperativeHandle(
    ref,
    () => {
      return { usedProcessors, setUsedProcessors };
    },
    [usedProcessors]
  );
  return (
    <>
      <AnimatePresence>
        {' '}
        {usedProcessors.length ? (
          usedProcessors.map((item, index) => (
            <Draggable key={item?.id} draggableId={item?.id} index={index}>
              {(provided) => (
                <div
                  ref={provided.innerRef}
                  {...provided.draggableProps}
                  {...provided.dragHandleProps}>
                  <motion.div
                    initial={{
                      opacity: 0,
                      y: -50,
                      scale: 0.2,
                      filter: 'blur(20px)',
                    }}
                    animate={{
                      opacity: 1,
                      scale: 1,
                      filter: 'blur(0px)',
                      y: 0,
                    }}
                    exit={{
                      y: -50,
                      opacity: 0,
                      scale: 0.8,
                      duration: 1,
                      filter: 'blur(20px)',
                    }}
                    transition={{
                      duration: 0.5,
                      scale: {
                        type: 'just',
                        stiffness: 500,
                        restDelta: 0.008,
                      },
                    }}>
                    <div
                      style={{
                        width: '100%',
                        padding: '20px 20px 20px 20px',
                        borderTopRightRadius: '10px',
                        borderTopLeftRadius: '10px',
                        background: '#000',
                        display: 'flex',
                        justifyItems: 'center',
                        justifyContent: 'space-between',
                        alignItems: 'center',
                      }}>
                      <p
                        style={{
                          fontSize: '16px',
                          fontWeight: 'bold',
                          color: 'white',
                        }}>
                        {item.name.split('').map((el, i) => (
                          <motion.span
                            initial={{ opacity: 0 }}
                            animate={{ opacity: 1 }}
                            transition={{
                              duration: 0.3,
                              delay: i / 5,
                            }}
                            key={i}>
                            {el}{' '}
                          </motion.span>
                        ))}
                      </p>
                      <CloseLarge
                        size={24}
                        style={{
                          color: 'white',
                          padding: '2px',
                        }}
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
                          borderBottomRightRadius: '10px',
                          borderBottomLeftRadius: '10px',
                          paddingTop: '20px',
                          paddingBottom: '10px',
                          background: '#f4f4f4',
                          marginBottom: '30px',
                        }}>
                        {item.arguments
                          ? item.arguments.map((arg, argIndex) => (
                              <div
                                key={arg?.id}
                                style={{
                                  marginBottom: '20px',
                                  marginLeft: '20px',
                                  marginRight: '20px',
                                }}>
                                {arg.type === 'Text' ? (
                                  <TextInput
                                    type="text"
                                    onChange={(e) => {
                                      setUsedProcessors((perv) => {
                                        perv.forEach((x) => {
                                          if (x.id === item.id) {
                                            x.arguments.forEach((y) => {
                                              if (y.id === arg.id) {
                                                y['value'] = e.target.value;
                                              }
                                            });
                                          }
                                        });

                                        return perv;
                                      });
                                      // arg.value = e.target.value;
                                      // item.arguments[argIndex] = arg;
                                      // usedProcessors[index] = item;
                                      // setUsedProcessors(
                                      //   (prevState) => usedProcessors
                                      // );
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
                                ) : arg.type === 'Boolean' ? (
                                  <Toggle
                                    id={arg.name + index}
                                    labelA="False"
                                    labelB="True"
                                    defaultToggled={arg.value}
                                    labelText={arg.name}
                                    onToggle={(props) => {
                                      arg.value = props;
                                      item.arguments[argIndex] = arg;
                                      usedProcessors[index] = item;
                                      setUsedProcessors(
                                        (prevState) => usedProcessors
                                      );
                                    }}
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
                  </motion.div>
                </div>
              )}
            </Draggable>
          ))
        ) : (
          <>
            <div
              style={{
                width: '100%',
                height: '100px',
                border: '2px dashed gray',
                borderColor: '#d2d2d2',
                borderRadius: '5px',
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center',
              }}>
              <p>Drop a processor here</p>
            </div>
          </>
        )}
      </AnimatePresence>
    </>
  );
});

export default UsedProcessorList;
