/* IMPORTS */
// import './components.js';
import { AddToast, GetLink, GetBoard } from './components.js';

/* MAIN FUNCTIONS */
window.onload = function() {
    var toast = document.getElementById('toast');    
    
    if (toast) {
        AddToast("Special Beam Canon'd to the server!", toast);
    }
}

document.addEventListener('DOMContentLoaded', function() {
    
    var form = document.querySelector('#drop');
    var fileInput = document.querySelector('#beam');
    var dataTransfer = new DataTransfer();
    var filenames = [];

    form.addEventListener('drop', function(e) {
        e.preventDefault();
        for (var i = 0; i < e.dataTransfer.files.length; i++) {
            dataTransfer.items.add(e.dataTransfer.files[i]);
            filenames.push(e.dataTransfer.files[i].name);
        }
        appendFilenames(filenames);
    });

    form.addEventListener('dragover', function(e) {
        e.preventDefault();
    });

    form.addEventListener('submit', function(e) {
        fileInput.files = dataTransfer.files;
    });
});

/* HELPER FUNCTIONS */
function appendFilenames(filenames) {
        
    let dropZone = GetBoard();

    // Removing the existing list if it exists
    let outdated = dropZone.querySelector('ul');
    if (outdated) {
        dropZone.removeChild(outdated);
    }

    var Beam = document.createElement('ul');
    for (var i = 0; i < filenames.length; i++) {
        var file = document.createElement('li');
        file.textContent = filenames[i];
        Beam.appendChild(file);
    }

    var button = dropZone.querySelector('button');
    dropZone.insertBefore(Beam, button);
}