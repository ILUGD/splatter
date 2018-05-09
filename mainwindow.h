#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>
#include "newform.h"

namespace Ui {
class MainWindow;
}

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    explicit MainWindow(QWidget *parent = 0);
    ~MainWindow();

    void createPoster();

private slots:
    void on_actionNew_triggered();

private:
    Ui::MainWindow *ui;
    NewForm *n;
};

#endif // MAINWINDOW_H
