'use client';

import React, { useEffect, useRef, useState } from 'react';
import { Content, Grid, Theme } from '@carbon/react';
import MainHeader from '@/components/MainHeader/MainHeader';
import PipelinePopup from '../../../components/CreatePipeline/PipelinePopup';
import PiplineNavbar from '../../../components/CreatePipeline/Navbar';
import { DragDropContext } from 'react-beautiful-dnd';
import UsedProcessors from '../../../components/CreatePipeline/UsedProcessors';
import Processors from '../../../components/CreatePipeline/Processors';
import { useRouter } from 'next/navigation';
import { getValueFromLocalStorage } from '../../../settings';
import uuidv4 from '../../../utils/uuidv4';

export default function EditPipelinePage() {
  const ProcessorsRef = useRef();
  const UsedProcessorsRef = useRef();
  const [pipelineName, setPipelineName] = useState('');
  const [errorMessageForSavePipeline, setErrorMessageForSavePipeline] =
    useState('');
  const router = useRouter();

  if (typeof window !== 'undefined')
    if (!localStorage.getItem('token')) window.location.assign('/auth');

  useEffect(() => {
    const pipeline = getValueFromLocalStorage('selected_pipeline');
    if (!pipeline) {
      return router.push('/pipelines');
    }

    const pj = Object.assign({}, { ...JSON.parse(pipeline) });
    if (pj) {
      setPipelineName(pj.name);
      const setUsedProcessors = UsedProcessorsRef.current.setUsedProcessors;
      debugger;

      if (pj.processors.length) {
        pj.processors.forEach((process) => {
          Object.assign(process, { id: uuidv4() });
          if (process.arguments?.length) {
            process.arguments.forEach((item) => {
              Object.assign(item, { id: uuidv4() });
            });
          }

          if (process?.output?.length) {
            process.output.forEach((item) => {
              Object.assign(item, { id: uuidv4() });
            });
          }
          if (process?.input?.length) {
            process.input.forEach((item) => {
              Object.assign(item, { id: uuidv4() });
            });
          }
        });
      }
      setUsedProcessors(pj.processors);
    }
  }, []);
  const onDragEnd = (result) => {
    const setUsedProcessors = UsedProcessorsRef.current.setUsedProcessors;
    const usedProcessors = [...UsedProcessorsRef.current.usedProcessors];

    const { destination, source, draggableId, type } = result;

    if (!destination) return;
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
      debugger;
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
  // const [errorMessage, setErrorMessage] = useState('');
  // const [errorMessageForSavePipeline, setErrorMessageForSavePipeline] =
  //   useState('');

  // const [usedProcessors, setUsedProcessors] = useState([]);

  // const [processors, setProcessors] = useState([]);

  // const [pipelineName, setPipelineName] = useState('');
  // const [openPipeline, setOpenPipeline] = useState(false);

  // const [searchTerm, setSearchTerm] = useState('');
  // const [searchResults, setSearchResults] = useState([]);

  // useEffect(() => {
  //   const results = processors.filter((listItem) =>
  //     listItem.name.toLowerCase().includes(searchTerm.toLowerCase()),
  //   );
  //   setSearchResults(results);
  // }, [searchTerm]);

  // useEffect(() => {
  //   const pipeline = getValueFromLocalStorage('selected_pipeline');
  //   if (!pipeline) {
  //     window.location.href = '/pipelines';
  //     return;
  //   }

  //   const pj = JSON.parse(pipeline);
  //   setPipelineName(pj.name);
  //   setUsedProcessors(pj.processors);
  //   setOpenPipeline(true);

  //   axios
  //     .get(Address() + '/api/v1/processors', {
  //       headers: {
  //         Authorization: `Bearer ${localStorage.getItem('token')}`,
  //       },
  //     })
  //     .then(function(response) {
  //       if (response.status !== 200) {
  //         setErrorMessage('Unexpected response from server');
  //         return;
  //       }
  //       setProcessors(response.data);
  //       setSearchResults(response.data);
  //     })
  //     .catch(function(error) {
  //       if (error.response) {
  //         if (error.response.status === 401) {
  //           window.location.href = '/auth';
  //           return;
  //         }
  //         setErrorMessage(error.response.data.message);
  //       } else {
  //         setErrorMessage('Unknown Error');
  //         console.error(error);
  //       }
  //     });
  // }, []);

  // const draggingItem = useRef();
  // const dragOverItem = useRef();

  // const handleDragStart = (e, position, source) => {
  //   draggingItem.current = position;
  //   draggingItem.source = source;
  // };

  // const handleDragEnter = (e, position, destination) => {
  //   dragOverItem.destination = destination;
  //   dragOverItem.current = position;
  // };

  // const handleDragEnd = (e) => {
  //   if (dragOverItem.destination === 'processors') {
  //     return;
  //   }

  //   const listCopy = [...usedProcessors];

  //   let draggingItemContent;
  //   if (draggingItem.source === 'processors') {
  //     draggingItemContent = JSON.parse(JSON.stringify(processors[draggingItem.current]));
  //     listCopy.splice(dragOverItem.current + 1, 0, draggingItemContent);
  //   } else {
  //     draggingItemContent = listCopy[draggingItem.current];
  //     listCopy.splice(draggingItem.current, 1);
  //     listCopy.splice(dragOverItem.current, 0, draggingItemContent);
  //   }

  //   draggingItem.current = null;
  //   dragOverItem.current = null;
  //   setUsedProcessors(listCopy);
  // };

  // const categoryToColor = (category) => {
  //   switch (category) {
  //     case 'Communication':
  //       return 'green';
  //     case 'Flow':
  //       return 'blue';
  //     case 'InputOutput':
  //       return 'red';
  //     default:
  //       return '';
  //   }
  // };

  // if (typeof window !== 'undefined')
  //   if (!localStorage.getItem('token')) window.location.assign('/auth');

  // return (
  //   <>
  //     <div>
  //       <Theme theme='g100'>
  //         <MainHeader />
  //       </Theme>
  //       <Content>
  //         {openPipeline ? (
  //           <Modal
  //             open
  //             preventCloseOnClickOutside={true}
  //             isFullWidth
  //             modalHeading='Edit/Duplicate pipeline'
  //             modalLabel='Pipeline information'
  //             primaryButtonText='Continue'
  //             secondaryButtonText='Back to pipelines'
  //             onRequestSubmit={(event) => {
  //               if (pipelineName) {
  //                 setOpenPipeline(false);
  //               }
  //             }}
  //             onRequestClose={(e) => {
  //               window.location.href = '/pipelines';
  //             }}
  //           >
  //             <div
  //               style={{
  //                 padding: '20px',
  //               }}
  //             >
  //               <TextInput
  //                 required={true}
  //                 data-modal-primary-focus
  //                 id='text-input-1'
  //                 labelText='Pipeline name'
  //                 placeholder='e.g. LOGIN_USER__V1'
  //                 style={{
  //                   marginBottom: '1rem',
  //                 }}
  //                 defaultValue={pipelineName}
  //                 onChange={(event) => setPipelineName(event.target.value)}
  //               />
  //             </div>
  //           </Modal>
  //         ) : null}

  //         <Grid className='create-page' fullWidth>
  //           <Column
  //             lg={16}
  //             md={8}
  //             sm={4}
  //             className='create-page_header'
  //             style={{ marginBottom: '48px' }}
  //           >
  //             <Breadcrumb>
  //               <BreadcrumbItem>
  //                 <a href='/'>Home</a>
  //               </BreadcrumbItem>
  //               <BreadcrumbItem>
  //                 <a href='/pipelines'>Pipelines</a>
  //               </BreadcrumbItem>
  //             </Breadcrumb>
  //             <Grid fullWidth>
  //               <Column md={4} lg={{ span: 7, offset: 0 }} sm={4}>
  //                 <h1 className='create-page__heading'>
  //                   Edit/Duplicate pipeline
  //                 </h1>
  //               </Column>
  //               <Column md={4} lg={{ span: 1, offset: 13 }} sm={4}>
  //                 <Button
  //                   onClick={(event) => {
  //                     axios
  //                       .post(
  //                         Address() + '/api/v1/pipelines',
  //                         {
  //                           name: pipelineName,
  //                           processors: usedProcessors,
  //                         },
  //                         {
  //                           headers: {
  //                             Authorization: `Bearer ${localStorage.getItem(
  //                               'token',
  //                             )}`,
  //                           },
  //                         },
  //                       )
  //                       .then(function(response) {
  //                         if (response.status !== 201) {
  //                           setErrorMessageForSavePipeline(
  //                             'Unexpected response from server',
  //                           );
  //                           return;
  //                         }
  //                         window.location.href = '/pipelines';
  //                       })
  //                       .catch(function(error) {
  //                         if (error.response) {
  //                           setErrorMessageForSavePipeline(
  //                             error.response.data.message,
  //                           );
  //                         } else {
  //                           setErrorMessageForSavePipeline('Unknown Error');
  //                         }
  //                       });
  //                   }}
  //                 >
  //                   Save {pipelineName}
  //                 </Button>
  //               </Column>
  //             </Grid>
  //           </Column>

  //           <Column md={4} lg={{ span: 7, offset: 1 }} sm={4}>
  //             {errorMessageForSavePipeline !== '' && (
  //               <InlineNotification
  //                 aria-label='closes notification'
  //                 kind='error'
  //                 statusIconDescription='notification'
  //                 subtitle={errorMessageForSavePipeline}
  //                 onClose={() => {
  //                   setErrorMessageForSavePipeline('');
  //                 }}
  //                 style={{ marginBottom: '16px', maxWidth: '500px' }}
  //               />
  //             )}
  //             <div
  //               style={{
  //                 paddingBottom: '30px',
  //                 paddingTop: '30px',
  //               }}
  //             >
  //               {usedProcessors.length ? (
  //                 usedProcessors.map((item, index) => (
  //                   <div
  //                     onDragStart={(e) =>
  //                       handleDragStart(e, index, 'used_processors')
  //                     }
  //                     onDragOver={(e) => e.preventDefault()}
  //                     onDragEnter={(e) =>
  //                       handleDragEnter(e, index, 'used_processors')
  //                     }
  //                     onDragEnd={handleDragEnd}
  //                     key={index}
  //                     draggable
  //                   >
  //                     <div
  //                       style={{
  //                         maxWidth: '400px',
  //                         padding: '20px 20px 20px 20px',
  //                         borderTopRightRadius: '10px',
  //                         borderTopLeftRadius: '10px',
  //                         background: '#000',
  //                         display: 'flex',
  //                         justifyItems: 'center',
  //                         justifyContent: 'space-between',
  //                         alignItems: 'center',
  //                       }}
  //                     >
  //                       <p
  //                         style={{
  //                           fontSize: '16px',
  //                           fontWeight: 'bold',
  //                           color: 'white',
  //                         }}
  //                       >
  //                         {item.name}
  //                       </p>
  //                       <CloseLarge
  //                         size={24}
  //                         style={{ color: 'white', padding: '2px' }}
  //                         onClick={(e) => {
  //                           setUsedProcessors(
  //                             usedProcessors.filter((value, index1) => {
  //                               return index1 !== index;
  //                             }),
  //                           );
  //                         }}
  //                       />
  //                     </div>
  //                     {item.arguments.length ? (
  //                       <div
  //                         style={{
  //                           maxWidth: '400px',
  //                           borderBottomRightRadius: '10px',
  //                           borderBottomLeftRadius: '10px',
  //                           paddingTop: '20px',
  //                           paddingBottom: '10px',
  //                           background: '#f4f4f4',
  //                           marginBottom: '30px',
  //                         }}
  //                       >
  //                         {item.arguments
  //                           ? item.arguments.map((arg, argIndex) => (
  //                             <div
  //                               key={arg.name}
  //                               style={{
  //                                 marginBottom: '20px',
  //                                 marginLeft: '20px',
  //                                 marginRight: '20px',
  //                               }}
  //                             >
  //                               {arg.type === 'Text' ? (
  //                                 <TextInput
  //                                   type='text'
  //                                   onChange={(e) => {
  //                                     arg.value = e.target.value;
  //                                     item.arguments[argIndex] = arg;
  //                                     usedProcessors[index] = item;
  //                                     setUsedProcessors(
  //                                       (prevState) => usedProcessors,
  //                                     );
  //                                   }}
  //                                   required={arg.required}
  //                                   placeholder={arg.name}
  //                                   defaultValue={arg.value}
  //                                 />
  //                               ) : arg.type === 'Numeric' ? (
  //                                 <TextInput
  //                                   type='text'
  //                                   onChange={(e) => {
  //                                     arg.value = parseInt(e.target.value);
  //                                     item.arguments[argIndex] = arg;
  //                                     usedProcessors[index] = item;
  //                                     setUsedProcessors(
  //                                       (prevState) => usedProcessors,
  //                                     );
  //                                   }}
  //                                   required={arg.required}
  //                                   defaultValue={arg.value}
  //                                   placeholder={arg.name + ' (Number)'}
  //                                 />
  //                               ) : arg.type === 'Boolean' ? (
  //                                 <Toggle
  //                                   id={arg.name + index}
  //                                   labelA='False' labelB='True'
  //                                   defaultToggled={arg.value}
  //                                   labelText={arg.name}
  //                                   onToggle={props => {
  //                                     arg.value = props;
  //                                     item.arguments[argIndex] = arg;
  //                                     usedProcessors[index] = item;
  //                                     setUsedProcessors(
  //                                       (prevState) => usedProcessors,
  //                                     );
  //                                   }}
  //                                 />
  //                               ) : (
  //                                 <p>
  //                                   No input found for argument {arg.name}
  //                                 </p>
  //                               )}

  //                               {/*{arg.required ? <p style={{ paddingTop: '3px', color: 'red' }}>*</p> : null}*/}
  //                               {/*<p style={{ fontWeight: 'bold', padding: '3px 10px 0 10px' }}>{arg.type}</p>*/}
  //                               {/*<p*/}
  //                               {/*  style={{ paddingTop: '3px' }}>{arg.name !== 'processable.body' ? arg.name : ''}</p>*/}
  //                             </div>
  //                           ))
  //                           : null}
  //                       </div>
  //                     ) : (
  //                       <p>No arguments</p>
  //                     )}
  //                   </div>
  //                 ))
  //               ) : (
  //                 <div
  //                   onDragOver={() => {
  //                     dragOverItem.destination = 'used_processors';
  //                   }}
  //                   style={{
  //                     maxWidth: '400px',
  //                     height: '100px',
  //                     border: '2px dashed gray',
  //                     borderColor: '#d2d2d2',
  //                     borderRadius: '5px',
  //                     display: 'flex',
  //                     alignItems: 'center',
  //                     justifyContent: 'center',
  //                   }}
  //                 >
  //                   <p>Drop a processor here</p>
  //                 </div>
  //               )}
  //             </div>
  //           </Column>

  //           <Column md={4} lg={{ span: 6, offset: 8 }} sm={4}>
  //             <StickyBox
  //               style={{ height: '20px' }}
  //               offsetTop={100}
  //               offsetBottom={20}
  //             >
  //               <ContainedList label='Processors' kind='on-page' action={''}>
  //                 <Search
  //                   placeholder='Filter'
  //                   closeButtonLabelText='Clear search input'
  //                   size='lg'
  //                   labelText='Filter search'
  //                   value={searchTerm}
  //                   onChange={(e) => {
  //                     setSearchTerm(e.target.value);
  //                   }}
  //                 />
  //                 {errorMessage !== '' && (
  //                   <InlineNotification
  //                     aria-label='closes notification'
  //                     kind='error'
  //                     statusIconDescription='notification'
  //                     subtitle={errorMessage}
  //                     onClose={() => {
  //                       setErrorMessage('');
  //                     }}
  //                     style={{ marginBottom: '16px' }}
  //                   />
  //                 )}
  //                 {searchResults &&
  //                   searchResults.map((item, index) => (
  //                     <div
  //                       onDragStart={(e) => {
  //                         handleDragStart(e, index, 'processors');
  //                       }}
  //                       onDragOver={(e) => e.preventDefault()}
  //                       onDragEnter={(e) =>
  //                         handleDragEnter(e, index, 'processors')
  //                       }
  //                       onDragEnd={handleDragEnd}
  //                       key={index}
  //                       draggable
  //                     >
  //                       <ExpandableTile
  //                         style={{
  //                           paddingLeft: '20px',
  //                           marginTop: '10px',
  //                           marginBottom: '10px',
  //                         }}
  //                         tileCollapsedIconText='Details'
  //                         tileExpandedIconText='Details'
  //                       >
  //                         <TileAboveTheFoldContent>
  //                           <div>
  //                             <h5>{item.name}</h5>
  //                             <p
  //                               style={{
  //                                 fontSize: '14px',
  //                               }}
  //                             >
  //                               {item.description}
  //                             </p>
  //                             <Tag type={categoryToColor(item.category)}>
  //                               {item.category}
  //                             </Tag>
  //                           </div>
  //                         </TileAboveTheFoldContent>
  //                         <TileBelowTheFoldContent>
  //                           <div style={{ marginTop: '10px' }}>
  //                             <p
  //                               style={{
  //                                 fontWeight: 'bold',
  //                                 fontSize: '18px',
  //                               }}
  //                             >
  //                               Arguments
  //                             </p>
  //                             {item.arguments ? (
  //                               item.arguments.map((arg, index) => (
  //                                 <div
  //                                   key={arg.name}
  //                                   style={{ display: 'flex' }}
  //                                 >
  //                                   {arg.required ? (
  //                                     <p
  //                                       style={{
  //                                         paddingTop: '3px',
  //                                         color: 'red',
  //                                       }}
  //                                     >
  //                                       *
  //                                     </p>
  //                                   ) : null}
  //                                   <p
  //                                     style={{
  //                                       fontWeight: 'bold',
  //                                       padding: '3px 10px 0 10px',
  //                                     }}
  //                                   >
  //                                     {arg.type}
  //                                   </p>
  //                                   <p style={{ paddingTop: '3px' }}>
  //                                     {arg.name !== 'processable.body'
  //                                       ? arg.name
  //                                       : ''}
  //                                   </p>
  //                                 </div>
  //                               ))
  //                             ) : (
  //                               <p>No arguments</p>
  //                             )}
  //                           </div>

  //                           <div style={{ marginTop: '10px' }}>
  //                             <p
  //                               style={{
  //                                 fontWeight: 'bold',
  //                                 fontSize: '18px',
  //                               }}
  //                             >
  //                               Input
  //                             </p>
  //                             {item.input ? (
  //                               item.input.map((arg, index) => (
  //                                 <div
  //                                   key={arg.name}
  //                                   style={{ display: 'flex' }}
  //                                 >
  //                                   {arg.required ? (
  //                                     <p
  //                                       style={{
  //                                         paddingTop: '3px',
  //                                         color: 'red',
  //                                       }}
  //                                     >
  //                                       *
  //                                     </p>
  //                                   ) : null}
  //                                   <p
  //                                     style={{
  //                                       fontWeight: 'bold',
  //                                       padding: '3px 10px 0 10px',
  //                                     }}
  //                                   >
  //                                     {arg.type}
  //                                   </p>
  //                                   <p style={{ paddingTop: '3px' }}>
  //                                     {arg.name !== 'processable.body'
  //                                       ? arg.name
  //                                       : ''}
  //                                   </p>
  //                                 </div>
  //                               ))
  //                             ) : (
  //                               <p>No input</p>
  //                             )}
  //                           </div>

  //                           <div style={{ marginTop: '10px' }}>
  //                             <p
  //                               style={{
  //                                 fontWeight: 'bold',
  //                                 fontSize: '18px',
  //                               }}
  //                             >
  //                               Output
  //                             </p>
  //                             {item.output ? (
  //                               item.output.map((arg, index) => (
  //                                 <div
  //                                   key={arg.name}
  //                                   style={{ display: 'flex' }}
  //                                 >
  //                                   {arg.required ? (
  //                                     <p
  //                                       style={{
  //                                         paddingTop: '3px',
  //                                         color: 'red',
  //                                       }}
  //                                     >
  //                                       *
  //                                     </p>
  //                                   ) : null}
  //                                   <p
  //                                     style={{
  //                                       fontWeight: 'bold',
  //                                       padding: '3px 10px 0 10px',
  //                                     }}
  //                                   >
  //                                     {arg.type}
  //                                   </p>
  //                                   <p style={{ paddingTop: '3px' }}>
  //                                     {arg.name !== 'processable.body'
  //                                       ? arg.name
  //                                       : ''}
  //                                   </p>
  //                                 </div>
  //                               ))
  //                             ) : (
  //                               <p>No output</p>
  //                             )}
  //                           </div>
  //                         </TileBelowTheFoldContent>
  //                       </ExpandableTile>
  //                     </div>
  //                   ))}
  //               </ContainedList>
  //             </StickyBox>
  //           </Column>
  //         </Grid>
  //       </Content>
  //     </div>
  //   </>
  // );
}
