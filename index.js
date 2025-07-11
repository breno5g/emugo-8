document.addEventListener("DOMContentLoaded", function () {
  const opcodeData = [
    {
      op: "00E0",
      mn: "CLS",
      cat: "display",
      desc: "Limpa a tela. Define todos os pixels do buffer gráfico para 0.",
    },
    {
      op: "00EE",
      mn: "RET",
      cat: "flow",
      desc: "Retorna de uma sub-rotina. Define o PC para o endereço no topo da pilha e decrementa o SP.",
    },
    {
      op: "1NNN",
      mn: "JP addr",
      cat: "flow",
      desc: "Salta para o endereço NNN. Define o PC para NNN.",
    },
    {
      op: "2NNN",
      mn: "CALL addr",
      cat: "flow",
      desc: "Chama a sub-rotina em NNN. Coloca o PC atual na pilha, incrementa SP e define PC para NNN.",
    },
    {
      op: "3XNN",
      mn: "SE Vx, byte",
      cat: "flow",
      desc: "Pula a próxima instrução se o registrador VX for igual ao valor NN.",
    },
    {
      op: "4XNN",
      mn: "SNE Vx, byte",
      cat: "flow",
      desc: "Pula a próxima instrução se o registrador VX NÃO for igual ao valor NN.",
    },
    {
      op: "5XY0",
      mn: "SE Vx, Vy",
      cat: "flow",
      desc: "Pula a próxima instrução se os registradores VX e VY forem iguais.",
    },
    {
      op: "6XNN",
      mn: "LD Vx, byte",
      cat: "memory",
      desc: "Carrega (Load) o valor NN no registrador VX.",
    },
    {
      op: "7XNN",
      mn: "ADD Vx, byte",
      cat: "math",
      desc: "Soma (Add) o valor NN ao registrador VX. O flag de carry VF não é afetado.",
    },
    {
      op: "8XY0",
      mn: "LD Vx, Vy",
      cat: "memory",
      desc: "Carrega (Load) o valor do registrador VY no registrador VX.",
    },
    {
      op: "8XY1",
      mn: "OR Vx, Vy",
      cat: "math",
      desc: "Realiza um OU (OR) bit a bit entre VX e VY, e armazena o resultado em VX.",
    },
    {
      op: "8XY2",
      mn: "AND Vx, Vy",
      cat: "math",
      desc: "Realiza um E (AND) bit a bit entre VX e VY, e armazena o resultado em VX.",
    },
    {
      op: "8XY3",
      mn: "XOR Vx, Vy",
      cat: "math",
      desc: "Realiza um OU Exclusivo (XOR) bit a bit entre VX e VY, e armazena o resultado em VX.",
    },
    {
      op: "8XY4",
      mn: "ADD Vx, Vy",
      cat: "math",
      desc: "Soma VY a VX. VF é definido como 1 se houver carry (resultado > 255), senão 0.",
    },
    {
      op: "8XY5",
      mn: "SUB Vx, Vy",
      cat: "math",
      desc: "Subtrai VY de VX. VF é 1 se NÃO houver borrow (VX >= VY), senão 0.",
    },
    {
      op: "8XY6",
      mn: "SHR Vx",
      cat: "math",
      desc: "Desloca (Shift Right) VX um bit para a direita. VF é definido com o valor do bit menos significativo antes do deslocamento.",
    },
    {
      op: "8XY7",
      mn: "SUBN Vx, Vy",
      cat: "math",
      desc: "Define VX como VY - VX. VF é 1 se NÃO houver borrow (VY >= VX), senão 0.",
    },
    {
      op: "8XYE",
      mn: "SHL Vx",
      cat: "math",
      desc: "Desloca (Shift Left) VX um bit para a esquerda. VF é definido com o valor do bit mais significativo antes do deslocamento.",
    },
    {
      op: "9XY0",
      mn: "SNE Vx, Vy",
      cat: "flow",
      desc: "Pula a próxima instrução se os registradores VX e VY NÃO forem iguais.",
    },
    {
      op: "ANNN",
      mn: "LD I, addr",
      cat: "memory",
      desc: "Carrega (Load) o endereço NNN no registrador de índice I.",
    },
    {
      op: "BNNN",
      mn: "JP V0, addr",
      cat: "flow",
      desc: "Salta para o endereço NNN + valor de V0.",
    },
    {
      op: "CXNN",
      mn: "RND Vx, byte",
      cat: "math",
      desc: "Gera um número aleatório (0-255), faz um AND com NN, e armazena em VX.",
    },
    {
      op: "DXYN",
      mn: "DRW Vx, Vy, nibble",
      cat: "display",
      desc: "Desenha um sprite de 8xN pixels em (VX, VY). Os dados do sprite são lidos da memória a partir do endereço I. VF é definido como 1 se houver colisão.",
    },
    {
      op: "EX9E",
      mn: "SKP Vx",
      cat: "input",
      desc: "Pula (Skip) a próxima instrução se a tecla com o valor de VX estiver pressionada.",
    },
    {
      op: "EXA1",
      mn: "SKNP Vx",
      cat: "input",
      desc: "Pula (Skip Not Pressed) a próxima instrução se a tecla com o valor de VX NÃO estiver pressionada.",
    },
    {
      op: "FX07",
      mn: "LD Vx, DT",
      cat: "timer",
      desc: "Carrega (Load) o valor do Delay Timer (DT) no registrador VX.",
    },
    {
      op: "FX0A",
      mn: "LD Vx, K",
      cat: "input",
      desc: "Pausa a execução até uma tecla ser pressionada. O valor da tecla é armazenado em VX.",
    },
    {
      op: "FX15",
      mn: "LD DT, Vx",
      cat: "timer",
      desc: "Define o Delay Timer (DT) com o valor de VX.",
    },
    {
      op: "FX18",
      mn: "LD ST, Vx",
      cat: "timer",
      desc: "Define o Sound Timer (ST) com o valor de VX.",
    },
    {
      op: "FX1E",
      mn: "ADD I, Vx",
      cat: "memory",
      desc: "Soma (Add) o valor de VX ao registrador de índice I.",
    },
    {
      op: "FX29",
      mn: "LD F, Vx",
      cat: "display",
      desc: "Define I para o local do sprite do dígito correspondente ao valor de VX.",
    },
    {
      op: "FX33",
      mn: "LD B, Vx",
      cat: "memory",
      desc: "Armazena a representação BCD (Decimal Codificado em Binário) de VX em memory[I], memory[I+1] e memory[I+2].",
    },
    {
      op: "FX55",
      mn: "LD [I], Vx",
      cat: "memory",
      desc: "Armazena os valores dos registradores V0 a VX na memória, começando no endereço I.",
    },
    {
      op: "FX65",
      mn: "LD Vx, [I]",
      cat: "memory",
      desc: "Lê valores da memória a partir do endereço I para os registradores V0 a VX.",
    },
  ];

  const tableBody = document.getElementById("opcode-table-body");
  const filterButtons = document.querySelectorAll(".filter-btn");

  function populateTable(filter = "all") {
    tableBody.innerHTML = "";
    const filteredData =
      filter === "all"
        ? opcodeData
        : opcodeData.filter((op) => op.cat === filter);

    filteredData.forEach((op) => {
      const row = document.createElement("tr");
      row.className = "opcode-row bg-white border-b cursor-pointer";
      row.innerHTML = `
                        <td class="px-6 py-4 font-mono font-medium text-gray-900">${op.op
        }</td>
                        <td class="px-6 py-4 font-mono">${op.mn}</td>
                        <td class="px-6 py-4 hidden md:table-cell">${op.desc.split(".")[0]
        }.</td>
                    `;

      const detailRow = document.createElement("tr");
      detailRow.className = "hidden bg-gray-50";
      detailRow.innerHTML = `<td colspan="3" class="px-6 py-4 text-gray-700">${op.desc}</td>`;

      row.addEventListener("click", () => {
        detailRow.classList.toggle("hidden");
      });

      tableBody.appendChild(row);
      tableBody.appendChild(detailRow);
    });
  }

  filterButtons.forEach((button) => {
    button.addEventListener("click", () => {
      const category = button.dataset.category;
      populateTable(category);
      filterButtons.forEach((btn) =>
        btn.classList.remove("bg-blue-600", "text-white")
      );
      filterButtons.forEach((btn) =>
        btn.classList.add("bg-gray-200", "text-gray-800")
      );
      button.classList.add("bg-blue-600", "text-white");
      button.classList.remove("bg-gray-200", "text-gray-800");
    });
  });

  const mobileMenuButton = document.getElementById("mobile-menu-button");
  const mobileMenu = document.getElementById("mobile-menu");
  mobileMenuButton.addEventListener("click", () => {
    mobileMenu.classList.toggle("hidden");
  });
  mobileMenu.addEventListener("click", (e) => {
    if (e.target.tagName === "A") {
      mobileMenu.classList.add("hidden");
    }
  });

  const sections = document.querySelectorAll("section");
  const navLinks = document.querySelectorAll(".nav-link");
  const header = document.getElementById("main-header");

  const observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          navLinks.forEach((link) => {
            link.classList.toggle(
              "active",
              link.getAttribute("href").substring(1) === entry.target.id
            );
          });
        }
      });
    },
    {
      rootMargin: `-${header.offsetHeight}px 0px 0px 0px`,
      threshold: 0.2,
    }
  );

  sections.forEach((section) => observer.observe(section));

  const ctx = document.getElementById("memoryChart").getContext("2d");
  new Chart(ctx, {
    type: "doughnut",
    data: {
      labels: [
        "ROM & RAM (3584 bytes)",
        "Interpretador/Fontes (512 bytes)",
      ],
      datasets: [
        {
          label: "Distribuição de Memória",
          data: [3584, 512],
          backgroundColor: [
            "rgba(37, 99, 235, 0.7)",
            "rgba(234, 179, 8, 0.7)",
          ],
          borderColor: ["rgba(37, 99, 235, 1)", "rgba(234, 179, 8, 1)"],
          borderWidth: 1,
        },
      ],
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        legend: {
          position: "top",
        },
        title: {
          display: true,
          text: "Distribuição da Memória CHIP-8",
        },
      },
    },
  });

  populateTable();
  document
    .querySelector('.filter-btn[data-category="all"]')
    .classList.add("bg-blue-600", "text-white");
  document
    .querySelector('.filter-btn[data-category="all"]')
    .classList.remove("bg-gray-200", "text-gray-800");
});
