import React from 'react';

const FileInfo = (props) => (
    <div className="file-info">
        <div className="desc"> 
            <span className="h">Ім'я : </span>
            <span className="t">{props.file.Name}</span>
        </div>
        <div className="desc"> 
            <span className="h">Розмір : </span>
            <span className="t">{props.file.Size}</span>
        </div>
        <div className="desc"> 
            <span className="h">Час останньої модифікації : </span>
            <span className="t">{props.file.ModTime} </span>
        </div>
        <div className="desc"> 
            <span className="h">Чи є папкою : </span>
            <span className="t">{props.file.IsDir + ''}</span>
        </div>
        <div className="desc"> 
            <div className="h">Попередній вигляд : </div>
            <div className="file-preview">{props.file.Preview}</div>
        </div>
    </div>
)

export default FileInfo;