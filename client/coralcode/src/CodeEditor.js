import React, { useState, useRef } from 'react';
import MonacoEditor from '@monaco-editor/react';
import coralLogo from './coral-logo.png';

const CodeEditor = () => {
  const [code, setCode] = useState('// Write your Coral code here');
  const [fileName, setFileName] = useState('untitled.clr');
  const [isRunning, setIsRunning] = useState(false);
  const [output, setOutput] = useState('');
  const editorRef = useRef(null); // To store Monaco editor instance

  const handleEditorChange = (value) => {
    setCode(value);
  };

  const handleEditorMount = (editor, monaco) => {
    editorRef.current = editor;
  };

  const saveFile = async () => {
    try {
      const payloadText = code; // <-- I’ll explain this below
      let payload = { filename: fileName, code: payloadText };
  
      const response = await fetch('http://localhost:8080/editor/save_new_file', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload),
      });
  
      if (!response.ok) throw new Error(`Server error: ${response.statusText}`);
  
      const result = await response.json();
      setOutput(JSON.stringify(result, null, 2));
      console.log('Save response:', result);
  
    } catch (error) {
      setOutput(`❌ Error: ${error.message}`);
      console.error('Error saving file:', error);
    } finally {
      setIsRunning(false);
    }
  
  

    const blob = new Blob([code], { type: 'text/plain' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = fileName;
    a.click();
    URL.revokeObjectURL(url);
  };

  const createNewFile = async () => {
    const confirmed = window.confirm('Create new file? Unsaved changes will be lost.');
    if (!confirmed) return;
  
    //saveFile(); // Save current file before creating a new one
    // Reset the editor
    setCode('// Write your Coral code here');
    setFileName('untitled.clr');
    setOutput('');
  
    // Notify the backend
    try {
      const payload = { filename: 'untitled.clr' }; // or whatever the default name is
      const response = await fetch('http://localhost:8080/editor/create_new_file', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload),
      });
  
      if (!response.ok) throw new Error(`Server error: ${response.statusText}`);
  
      const result = await response.json();
      console.log('Create new file response:', result);
  
    } catch (error) {
      console.error('Error creating new file:', error);
      setOutput(`❌ Error creating new file: ${error.message}`);
    }
  };

  const runCode = async () => {
    if (!editorRef.current) {
      console.error('Editor not initialized');
      return;
    }

    setIsRunning(true);
    setOutput('');

    try {
      const selection = editorRef.current
        .getModel()
        .getValueInRange(editorRef.current.getSelection());
      const payloadText = selection.length > 0 ? selection : code;

      let payload;
      try {
        payload = JSON.parse(payloadText);
      } catch {
        payload = { code: payloadText };
      }

      const response = await fetch('http://localhost:8080/editor/compile', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload),
      });

      if (!response.ok) throw new Error(`Server error: ${response.statusText}`);

      const result = await response.json();
      setOutput(JSON.stringify(result, null, 2)); // ✅ Show result
      console.log('Run response:', result);

    } catch (error) {
      setOutput(`❌ Error: ${error.message}`);
      console.error('Error running code:', error);
    } finally {
      setIsRunning(false);
    }
  };

  return (
    <div className="editor-container">
      <div className="editor-header">
        <div className="left-controls">
          <input
            type="text"
            value={fileName}
            onChange={(e) => setFileName(e.target.value)}
            placeholder="Enter file name"
          />
          <button onClick={saveFile}>Save File</button>
          <button onClick={createNewFile}>New File</button>
          <button onClick={runCode} disabled={isRunning}>
            {isRunning ? 'Running...' : 'Run Code'}
          </button>
        </div>

        <div className="logo-heading">
          <img src={coralLogo} alt="Coral Logo" className="coral-logo" />
          <h1 className="coral-heading">CoralCode</h1>
        </div>
      </div>

      <MonacoEditor
        height="60vh"
        language="javascript"
        value={code}
        theme="vs-dark"
        onChange={handleEditorChange}
        onMount={handleEditorMount}
        options={{
          fontSize: 14,
          automaticLayout: true,
        }}
      />

      {/* ✅ Output section */}
      <div className="output-container">
        <h2>Output</h2>
        <pre>{output}</pre>
      </div>
    </div>
  );
};

export default CodeEditor;
